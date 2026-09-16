package repository_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/repository"
)

func TestSQLiteRepository(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "sqlite-repo-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	dbPath := filepath.Join(tmpDir, "test.db")
	repo, err := repository.NewSQLiteRepository(dbPath)
	if err != nil {
		t.Fatalf("failed to create sqlite repo: %v", err)
	}
	defer repo.Close()

	ctx := context.Background()

	// 1. Save and Get Problem
	p := &model.Problem{
		ID:          "two_sum",
		Title:       "Two Sum",
		Description: "Find two indices that sum to target",
		Language:    "python",
		Entrypoint:  "solve",
		Tests: []model.TestCase{
			{
				ID:             "1",
				Input:          "{\"nums\": [2, 7, 11, 15], \"target\": 9}",
				ExpectedOutput: "[0, 1]",
			},
		},
		AcceptedStrategies: []string{"hashmap_two_sum"},
		RequiredConcepts:   []string{"hashmap"},
		OptionalConcepts:   []string{"iteration"},
		PrimaryConcepts:    []string{"hashmap"},
	}

	if err := repo.SaveProblem(ctx, p); err != nil {
		t.Fatalf("failed to save problem: %v", err)
	}

	gotP, err := repo.GetProblem(ctx, "two_sum")
	if err != nil {
		t.Fatalf("failed to get problem: %v", err)
	}
	if gotP.Title != p.Title || len(gotP.Tests) != 1 || gotP.RequiredConcepts[0] != "hashmap" {
		t.Errorf("retrieved problem mismatch: %+v", gotP)
	}

	list, err := repo.ListProblems(ctx)
	if err != nil {
		t.Fatalf("failed to list problems: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 problem in list, got %d", len(list))
	}

	// 2. Save and Get Evaluation
	eval := &model.Evaluation{
		ID:             "eval-123",
		ProblemID:      "two_sum",
		SubmissionCode: "def solve(): pass",
		Explanation:    "using hashmap",
		TestResult:     model.TestStatusPass,
		PassedTests:    1,
		FailedTests:    0,
		TotalTests:     1,
		FidelityScore:  0.98,
		Decision:       model.DecisionAccept,
		CreatedAt:      time.Now().Truncate(time.Second),
	}

	if err := repo.SaveEvaluation(ctx, eval); err != nil {
		t.Fatalf("failed to save evaluation: %v", err)
	}

	gotEval, err := repo.GetEvaluation(ctx, "eval-123")
	if err != nil {
		t.Fatalf("failed to get evaluation: %v", err)
	}
	if gotEval.ID != eval.ID || gotEval.Decision != model.DecisionAccept || gotEval.FidelityScore != 0.98 {
		t.Errorf("retrieved evaluation mismatch: %+v", gotEval)
	}
}
