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
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/runner"
)

func setupEvaluator(t *testing.T) *evaluator.Evaluator {
	repoPath := filepath.Join("..", "..", "data", "dsa")
	repo, err := dsa.NewRepository(repoPath)
	if err != nil {
		t.Fatalf("failed to load ontology: %v", err)
	}

	embedder := nlp.NewLocalHashingEmbedder(256)
	ctx := context.Background()

	matcher, err := nlp.NewCascadeMatcher(ctx, repo.Ontology(), embedder, true)
	if err != nil {
		t.Fatalf("failed to init cascade matcher: %v", err)
	}

	r := runner.NewLocalRunner(2000)
	scriptPath := filepath.Join("..", "analysis", "python_ast.py")
	a := analysis.NewPythonAnalyzer(scriptPath)
	s := fidelity.NewScorer(repo.Ontology(), fidelity.DefaultThresholds())

	return evaluator.NewEvaluator(r, a, matcher, s, nil)
}

func TestEvaluatorPipeline(t *testing.T) {
	eval := setupEvaluator(t)
	ctx := context.Background()

	problem := model.Problem{
		ID:         "two_sum",
		Title:      "Two Sum",
		Language:   "python",
		Entrypoint: "solve",
		Tests: []model.TestCase{
			{
				ID:             "1",
				Input:          "{\"nums\": [2, 7, 11, 15], \"target\": 9}",
				ExpectedOutput: "[0, 1]",
			},
			{
				ID:             "2",
				Input:          "{\"nums\": [3, 2, 4], \"target\": 6}",
				ExpectedOutput: "[1, 2]",
			},
		},
		AcceptedStrategies: []string{"hashmap_two_sum"},
		RequiredConcepts:   []string{"hashmap"},
		PrimaryConcepts:    []string{"hashmap"},
	}

	// 1. Failed Test Case -> Immediate Rejection
	t.Run("failed test case stops immediately", func(t *testing.T) {
		sub := model.Submission{
			ProblemID:   "two_sum",
			SourceCode:  "def solve(nums, target): return [99, 99]",
			Explanation: "I'll store each number in a hashmap.",
		}

		res, err := eval.Evaluate(ctx, sub, problem)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if res.TestResult != model.TestStatusFail {
			t.Errorf("expected TestResult FAIL, got %s", res.TestResult)
		}
		if res.Decision != model.DecisionReject {
			t.Errorf("expected Decision REJECT, got %s", res.Decision)
		}
		if len(res.ActualConcepts) != 0 {
			t.Errorf("expected 0 actual concepts on test failure, got %d", len(res.ActualConcepts))
		}
	})

	// 2. Passing Test Case & Matching Approach -> ACCEPT
	t.Run("passing test and matching approach -> ACCEPT", func(t *testing.T) {
		code := `
def solve(nums, target):
    seen = {}
    for i, n in enumerate(nums):
        diff = target - n
        if diff in seen:
            return [seen[diff], i]
        seen[n] = i
    return []
`
		sub := model.Submission{
			ProblemID:   "two_sum",
			SourceCode:  code,
			Explanation: "I will use a hashmap to store the seen values and their indices, and look up the complement in O(1).",
		}

		res, err := eval.Evaluate(ctx, sub, problem)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if res.TestResult != model.TestStatusPass {
			t.Fatalf("expected TestResult PASS, got %s", res.TestResult)
		}
		if res.Decision != model.DecisionAccept {
			t.Errorf("expected Decision ACCEPT, got %s (score: %f)", res.Decision, res.FidelityScore)
		}
		if res.FidelityScore < 0.95 {
			t.Errorf("expected score >= 0.95, got %f", res.FidelityScore)
		}
	})

	// 3. Passing Test Case but Contradicting Claimed Approach -> REJECT
	t.Run("passing test but claiming different approach -> REJECT", func(t *testing.T) {
		code := `
def solve(nums, target):
    seen = {}
    for i, n in enumerate(nums):
        diff = target - n
        if diff in seen:
            return [seen[diff], i]
        seen[n] = i
    return []
`
		sub := model.Submission{
			ProblemID:   "two_sum",
			SourceCode:  code,
			Explanation: "I will use binary search to locate the elements.",
		}

		res, err := eval.Evaluate(ctx, sub, problem)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if res.TestResult != model.TestStatusPass {
			t.Fatalf("expected PASS, got %s", res.TestResult)
		}
		if res.Decision != model.DecisionReject {
			t.Errorf("expected Decision REJECT for claimed binary search, got %s (score: %f)", res.Decision, res.FidelityScore)
		}
	})

	// 4. User's exact submission (with top-level print before def and twoSum naming) -> ACCEPT
	t.Run("user exact submission with forward print and twoSum name -> ACCEPT", func(t *testing.T) {
		userCode := `nums = [2, 7, 11, 15]
target = 9

print(twoSum(nums, target))  # [0, 1]

def twoSum(nums, target):
    seen = {}

    for i, num in enumerate(nums):
        complement = target - num

        if complement in seen:
            return [seen[complement], i]

        seen[num] = i

    return []`

		userExplanation := "So I will iterate through the array, keep a Hashmap storing their index, match target - current number, check for this value in hte Hashmap, if found return these two index the current and the one stored in Hashmap else return empty array for not found."

		sub := model.Submission{
			ProblemID:   "two_sum",
			SourceCode:  userCode,
			Explanation: userExplanation,
		}

		userProblem := problem
		userProblem.EntrypointAliases = []string{"twoSum", "two_sum"}

		res, err := eval.Evaluate(ctx, sub, userProblem)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if res.TestResult != model.TestStatusPass {
			t.Fatalf("expected TestResult PASS, got %s (diags: %v)", res.TestResult, res.Diagnostics)
		}
		if res.Decision != model.DecisionAccept {
			t.Errorf("expected Decision ACCEPT, got %s (score: %f, diags: %v)", res.Decision, res.FidelityScore, res.Diagnostics)
		}
		if res.FidelityScore < 0.95 {
			t.Errorf("expected score >= 0.95, got %f", res.FidelityScore)
		}
	})
}
