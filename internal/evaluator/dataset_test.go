package evaluator_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/analysis"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/dsa"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/evaluator"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/fidelity"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/nlp"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/repository"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/runner"
)

func TestDataset2KEvaluationPipeline(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "eval_dataset.db")

	repo, err := repository.NewSQLiteRepository(dbPath)
	if err != nil {
		t.Fatalf("failed to init repository: %v", err)
	}
	defer repo.Close()

	problemsDir := filepath.Join("..", "..", "data", "problems")
	ctx := context.Background()

	if err := repository.LoadProblemsFromDirectory(ctx, problemsDir, repo); err != nil {
		t.Fatalf("failed to load dataset: %v", err)
	}

	allProblems, err := repo.ListProblems(ctx)
	if err != nil {
		t.Fatalf("failed to list problems: %v", err)
	}
	if len(allProblems) < 3600 {
		t.Fatalf("expected >= 3600 problems loaded, got %d", len(allProblems))
	}

	// Verify broad coverage across core algorithmic paradigms in dataset
	topicPresence := make(map[string]bool)
	for _, p := range allProblems {
		for _, tag := range p.TopicTags {
			topicPresence[tag] = true
		}
	}
	expectedTags := []string{
		"array", "dynamic-programming", "string", "hash-table",
		"binary-search", "greedy", "depth-first-search", "breadth-first-search",
		"tree", "sorting", "two-pointers", "backtracking",
	}
	for _, tag := range expectedTags {
		if !topicPresence[tag] {
			t.Errorf("expected tag '%s' to be present in ingested 2k dataset", tag)
		}
	}

	// Setup evaluator
	dsaRepoPath := filepath.Join("..", "..", "data", "dsa")
	dsaRepo, err := dsa.NewRepository(dsaRepoPath)
	if err != nil {
		t.Fatalf("failed to load ontology: %v", err)
	}

	embedder := nlp.NewLocalHashingEmbedder(256)
	matcher, err := nlp.NewCascadeMatcher(ctx, dsaRepo.Ontology(), embedder, true)
	if err != nil {
		t.Fatalf("failed to init matcher: %v", err)
	}

	r := runner.NewLocalRunner(3000)
	scriptPath := filepath.Join("..", "analysis", "python_ast.py")
	analyzer := analysis.NewPythonAnalyzer(scriptPath)
	scorer := fidelity.NewScorer(dsaRepo.Ontology(), fidelity.DefaultThresholds())
	eval := evaluator.NewEvaluator(r, analyzer, matcher, scorer, nil)

	// Fetch lc_1_two_sum from SQLite repository
	prob, err := repo.GetProblem(ctx, "lc_1_two_sum")
	if err != nil {
		t.Fatalf("failed to find lc_1_two_sum: %v", err)
	}

	// 1. Valid Solution Matching Code & Explanation
	t.Run("valid submission on 2k dataset problem", func(t *testing.T) {
		sub := model.Submission{
			ProblemID: prob.ID,
			SourceCode: `
def twoSum(nums, target=0):
    seen = {}
    for i, x in enumerate(nums):
        if x in seen:
            return seen[x]
        seen[x] = 0
    return seen.get(0, 0)
`,
			Explanation: "We use an array iteration and a hashmap to record numbers and find matching pairs.",
		}

		res, err := eval.Evaluate(ctx, sub, *prob)
		if err != nil {
			t.Fatalf("evaluation failed: %v", err)
		}
		t.Logf("Valid evaluation: Score=%.3f, Decision=%s, Matched=%v, Diags=%v",
			res.FidelityScore, res.Decision, res.MatchedConcepts, res.Diagnostics)
		if res.TestResult != model.TestStatusPass {
			t.Errorf("expected tests to PASS, got %s", res.TestResult)
		}
		if res.Decision != model.DecisionAccept {
			t.Errorf("expected ACCEPT decision, got %s (score: %.3f)", res.Decision, res.FidelityScore)
		}
		if res.FidelityScore < 0.95 {
			t.Errorf("expected high fidelity score >= 0.95, got %.3f", res.FidelityScore)
		}
	})

	// 2. Failing Solution -> Immediate Rejection
	t.Run("failing submission on 2k dataset problem", func(t *testing.T) {
		sub := model.Submission{
			ProblemID: prob.ID,
			SourceCode: `
def twoSum(nums, target=0):
    return -9999
`,
			Explanation: "We iterate through the array.",
		}

		res, err := eval.Evaluate(ctx, sub, *prob)
		if err != nil {
			t.Fatalf("evaluation failed: %v", err)
		}
		if res.TestResult != model.TestStatusFail {
			t.Errorf("expected test FAIL, got %s", res.TestResult)
		}
		if res.Decision != model.DecisionReject {
			t.Errorf("expected REJECT decision, got %s", res.Decision)
		}
	})

	// 3. Bluffing Explanation -> Fidelity Divergence Penalty
	t.Run("bluffing explanation penalized on 2k dataset problem", func(t *testing.T) {
		sub := model.Submission{
			ProblemID: prob.ID,
			SourceCode: `
def twoSum(nums, target=0):
    seen = {}
    for i, x in enumerate(nums):
        seen[x] = 0
    return seen.get(0, 0)
`,
			Explanation: "We construct a segment tree with lazy propagation and run Dijkstra shortest path algorithm.",
		}

		res, err := eval.Evaluate(ctx, sub, *prob)
		if err != nil {
			t.Fatalf("evaluation failed: %v", err)
		}
		if res.TestResult != model.TestStatusPass {
			t.Errorf("expected test PASS, got %s", res.TestResult)
		}
		// Code has no segment tree or dijkstra -> code/explanation divergence
		if res.Decision == model.DecisionAccept {
			t.Errorf("expected penalized score for bluffing explanation, got score=%.3f, decision=%s", res.FidelityScore, res.Decision)
		}
	})

	// 4. LeetCode Accounts Merge Solution with top-level self signature
	t.Run("leetcode accounts merge with top level self passes tests and gets accepted", func(t *testing.T) {
		advProb, err := repo.GetProblem(ctx, "adv_accounts_merge_union_find")
		if err != nil {
			t.Fatalf("failed to find adv_accounts_merge_union_find: %v", err)
		}

		sub := model.Submission{
			ProblemID: advProb.ID,
			SourceCode: `
def accountsMerge(self, accounts):
    parent = {}
    rank = {}
    email_to_name = {}

    def find(email):
        if parent[email] != email:
            parent[email] = find(parent[email])
        return parent[email]

    def union(email1, email2):
        root1 = find(email1)
        root2 = find(email2)
        if root1 == root2:
            return
        if rank[root1] < rank[root2]:
            root1, root2 = root2, root1
        parent[root2] = root1
        if rank[root1] == rank[root2]:
            rank[root1] += 1

    for account in accounts:
        name = account[0]
        first_email = account[1]
        if first_email not in parent:
            parent[first_email] = first_email
            rank[first_email] = 0
        email_to_name[first_email] = name

        for email in account[2:]:
            if email not in parent:
                parent[email] = email
                rank[email] = 0
            email_to_name[email] = name
            union(first_email, email)

    groups = {}
    for email in parent:
        root = find(email)
        if root not in groups:
            groups[root] = []
        groups[root].append(email)

    result = []
    for emails in groups.values():
        emails.sort()
        result.append([email_to_name[emails[0]]] + emails)
    return result
`,
			Explanation: "I used a disjoint set union-find data structure with path compression and rank optimization to merge connected accounts sharing emails.",
		}

		res, err := eval.Evaluate(ctx, sub, *advProb)
		if err != nil {
			t.Fatalf("evaluation failed: %v", err)
		}
		if res.TestResult != model.TestStatusPass {
			t.Errorf("expected test PASS, got %s (diags: %v)", res.TestResult, res.Diagnostics)
		}
		if res.Decision != model.DecisionAccept {
			t.Errorf("expected ACCEPT decision, got %s (score: %.3f, diags: %v)", res.Decision, res.FidelityScore, res.Diagnostics)
		}
	})

	// 5. LeetCode Accounts Merge Solution inside class Solution
	t.Run("leetcode accounts merge inside class Solution passes tests and gets accepted", func(t *testing.T) {
		advProb, err := repo.GetProblem(ctx, "adv_accounts_merge_union_find")
		if err != nil {
			t.Fatalf("failed to find adv_accounts_merge_union_find: %v", err)
		}

		classCode := `
class Solution:
    def accountsMerge(self, accounts: list[list[str]]) -> list[list[str]]:
        parent = {}
        rank = {}
        email_to_name = {}

        def find(email):
            if parent[email] != email:
                parent[email] = find(parent[email])
            return parent[email]

        def union(email1, email2):
            root1 = find(email1)
            root2 = find(email2)
            if root1 == root2:
                return
            if rank[root1] < rank[root2]:
                root1, root2 = root2, root1
            parent[root2] = root1
            if rank[root1] == rank[root2]:
                rank[root1] += 1

        for account in accounts:
            name = account[0]
            first_email = account[1]
            if first_email not in parent:
                parent[first_email] = first_email
                rank[first_email] = 0
            email_to_name[first_email] = name

            for email in account[2:]:
                if email not in parent:
                    parent[email] = email
                    rank[email] = 0
                email_to_name[email] = name
                union(first_email, email)

        groups = {}
        for email in parent:
            root = find(email)
            if root not in groups:
                groups[root] = []
            groups[root].append(email)

        result = []
        for emails in groups.values():
            emails.sort()
            result.append([email_to_name[emails[0]]] + emails)
        return result
`
		sub := model.Submission{
			ProblemID:   advProb.ID,
			SourceCode:  classCode,
			Explanation: "I implemented union find disjoint sets with rank heuristic and path compression to merge accounts.",
		}

		res, err := eval.Evaluate(ctx, sub, *advProb)
		if err != nil {
			t.Fatalf("evaluation failed: %v", err)
		}
		if res.TestResult != model.TestStatusPass {
			t.Errorf("expected test PASS, got %s (diags: %v)", res.TestResult, res.Diagnostics)
		}
		if res.Decision != model.DecisionAccept {
			t.Errorf("expected ACCEPT decision, got %s (score: %.3f, diags: %v)", res.Decision, res.FidelityScore, res.Diagnostics)
		}
	})
}
