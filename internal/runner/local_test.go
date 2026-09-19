package runner_test

import (
	"context"
	"strings"
	"testing"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/runner"
)

func TestLocalRunner(t *testing.T) {
	r := runner.NewLocalRunner(2000)
	ctx := context.Background()

	problem := model.Problem{
		ID:         "two_sum",
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
	}

	// Case 1: Valid Solution
	validCode := `
def solve(nums, target):
    seen = {}
    for i, n in enumerate(nums):
        diff = target - n
        if diff in seen:
            return [seen[diff], i]
        seen[n] = i
    return []
`
	res, err := r.Run(ctx, model.Submission{SourceCode: validCode}, problem)
	if err != nil {
		t.Fatalf("unexpected error running tests: %v", err)
	}
	if res.Status != model.TestStatusPass {
		t.Errorf("expected PASS, got %s (error: %s)", res.Status, res.Error)
	}
	if res.PassedCount != 2 || res.FailedCount != 0 {
		t.Errorf("expected 2 passed 0 failed, got passed=%d failed=%d", res.PassedCount, res.FailedCount)
	}

	// Case 2: Incorrect Solution
	wrongCode := `
def solve(nums, target):
    return [0, 0]
`
	res, err = r.Run(ctx, model.Submission{SourceCode: wrongCode}, problem)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Status != model.TestStatusFail {
		t.Errorf("expected FAIL, got %s", res.Status)
	}
	if res.FailedCount == 0 {
		t.Errorf("expected at least 1 failed test, got 0")
	}

	// Case 3: Syntax Error
	badSyntax := `def solve(nums, target) return`
	res, err = r.Run(ctx, model.Submission{SourceCode: badSyntax}, problem)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Status != model.TestStatusFail {
		t.Errorf("expected FAIL for syntax error, got %s", res.Status)
	}

	// Case 4: Infinite loop / Timeout
	timeoutRunner := runner.NewLocalRunner(500)
	infLoop := `
def solve(nums, target):
    while True:
        pass
`
	res, err = timeoutRunner.Run(ctx, model.Submission{SourceCode: infLoop}, problem)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Status != model.TestStatusFail {
		t.Errorf("expected timeout FAIL, got %s", res.Status)
	}

	// Case 5: LeetCode class Solution
	solutionClassCode := `
class Solution:
    def solve(self, nums, target):
        seen = {}
        for i, n in enumerate(nums):
            diff = target - n
            if diff in seen:
                return [seen[diff], i]
            seen[n] = i
        return []
`
	res, err = r.Run(ctx, model.Submission{SourceCode: solutionClassCode}, problem)
	if err != nil {
		t.Fatalf("unexpected error running class Solution: %v", err)
	}
	if res.Status != model.TestStatusPass {
		t.Errorf("expected class Solution to PASS, got %s (error: %s)", res.Status, res.Error)
	}

	// Case 6: Single-parameter function receiving list input
	singleArgProb := model.Problem{
		ID:         "second_max",
		Language:   "python",
		Entrypoint: "solve",
		Tests: []model.TestCase{
			{
				ID:             "1",
				Input:          "[10, 20, 4, 45, 99]",
				ExpectedOutput: "45",
			},
		},
	}
	singleArgCode := `
def solve(nums):
    unique = sorted(list(set(nums)))
    return unique[-2]
`
	res, err = r.Run(ctx, model.Submission{SourceCode: singleArgCode}, singleArgProb)
	if err != nil {
		t.Fatalf("unexpected error running single list argument: %v", err)
	}
	if res.Status != model.TestStatusPass {
		t.Errorf("expected single list argument to PASS, got %s (error: %s)", res.Status, res.Error)
	}

	// Case 7: Positional argument mapping when dict keys differ (e.g. solve(a, b) with {"nums": [...], "target": 9})
	positionalCode := `
def solve(a, b):
    seen = {}
    for i, n in enumerate(a):
        diff = b - n
        if diff in seen:
            return [seen[diff], i]
        seen[n] = i
    return []
`
	res, err = r.Run(ctx, model.Submission{SourceCode: positionalCode}, problem)
	if err != nil {
		t.Fatalf("unexpected error running positional mapping: %v", err)
	}
	if res.Status != model.TestStatusPass {
		t.Errorf("expected positional mapping to PASS, got %s (error: %s)", res.Status, res.Error)
	}

	// Case 8: Function accepting subset of dictionary keys (e.g. solve(nums) for {"nums": [...], "target": 9})
	singleParamDictProb := model.Problem{
		ID:         "two_sum_single",
		Language:   "python",
		Entrypoint: "solve",
		Tests: []model.TestCase{
			{
				ID:             "1",
				Input:          "{\"nums\": [1, 2, 3], \"target\": 3}",
				ExpectedOutput: "3",
			},
		},
	}
	singleParamDictCode := `
def solve(nums):
    return len(nums)
`
	res, err = r.Run(ctx, model.Submission{SourceCode: singleParamDictCode}, singleParamDictProb)
	if err != nil {
		t.Fatalf("unexpected error running partial dict param: %v", err)
	}
	if res.Status != model.TestStatusPass {
		t.Errorf("expected partial dict param to PASS, got %s (error: %s)", res.Status, res.Error)
	}

	// Case 9: LeetCode-style top-level method with self parameter (pasted from LeetCode editor without class Solution)
	accountsProb := model.Problem{
		ID:         "accounts_merge",
		Language:   "python",
		Entrypoint: "accountsMerge",
		Tests: []model.TestCase{
			{
				ID:             "1",
				Input:          "{\"accounts\": [[\"John\", \"johnsmith@mail.com\", \"john_newyork@mail.com\"], [\"John\", \"johnsmith@mail.com\", \"john00@mail.com\"], [\"Mary\", \"mary@mail.com\"], [\"John\", \"johnnybravo@mail.com\"]]}",
				ExpectedOutput: "[[\"John\", \"john00@mail.com\", \"john_newyork@mail.com\", \"johnsmith@mail.com\"], [\"Mary\", \"mary@mail.com\"], [\"John\", \"johnnybravo@mail.com\"]]",
			},
			{
				ID:             "2",
				Input:          "{\"accounts\": [[\"Gabe\", \"Gabe0@m.co\", \"Gabe3@m.co\", \"Gabe1@m.co\"], [\"Kevin\", \"Kevin3@m.co\", \"Kevin5@m.co\", \"Kevin0@m.co\"], [\"Ethan\", \"Ethan5@m.co\", \"Ethan4@m.co\", \"Ethan0@m.co\"], [\"Hanzo\", \"Hanzo3@m.co\", \"Hanzo1@m.co\", \"Hanzo0@m.co\"], [\"Fern\", \"Fern5@m.co\", \"Fern1@m.co\", \"Fern0@m.co\"]]}",
				ExpectedOutput: "[[\"Ethan\", \"Ethan0@m.co\", \"Ethan4@m.co\", \"Ethan5@m.co\"], [\"Fern\", \"Fern0@m.co\", \"Fern1@m.co\", \"Fern5@m.co\"], [\"Gabe\", \"Gabe0@m.co\", \"Gabe1@m.co\", \"Gabe3@m.co\"], [\"Hanzo\", \"Hanzo0@m.co\", \"Hanzo1@m.co\", \"Hanzo3@m.co\"], [\"Kevin\", \"Kevin0@m.co\", \"Kevin3@m.co\", \"Kevin5@m.co\"]]",
			},
		},
	}
	topLevelSelfCode := `
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
`
	res, err = r.Run(ctx, model.Submission{SourceCode: topLevelSelfCode}, accountsProb)
	if err != nil {
		t.Fatalf("unexpected error running top-level self accountsMerge: %v", err)
	}
	if res.Status != model.TestStatusPass {
		t.Errorf("expected top-level self accountsMerge to PASS, got %s (error: %s)", res.Status, res.Error)
	}

	// Case 10: LeetCode class Solution wrapping the exact same accountsMerge
	var classSolutionCode string
	for _, line := range strings.Split(strings.TrimSpace(topLevelSelfCode), "\n") {
		classSolutionCode += "    " + line + "\n"
	}
	classSolutionCode = "class Solution:\n" + classSolutionCode
	res, err = r.Run(ctx, model.Submission{SourceCode: classSolutionCode}, accountsProb)
	if err != nil {
		t.Fatalf("unexpected error running class Solution accountsMerge: %v", err)
	}
	if res.Status != model.TestStatusPass {
		t.Errorf("expected class Solution accountsMerge to PASS, got %s (error: %s)", res.Status, res.Error)
	}
}
