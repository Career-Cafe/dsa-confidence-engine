package repository_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/repository"
)

func TestLoadProblemsFromDirectory(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "prob-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	dbPath := filepath.Join(tmpDir, "test.db")
	repo, err := repository.NewSQLiteRepository(dbPath)
	if err != nil {
		t.Fatalf("failed to init repo: %v", err)
	}
	defer repo.Close()

	problemsDir := filepath.Join("..", "..", "data", "problems")
	ctx := context.Background()
	if err := repository.LoadProblemsFromDirectory(ctx, problemsDir, repo); err != nil {
		t.Fatalf("failed to load problems from directory: %v", err)
	}

	problems, err := repo.ListProblems(ctx)
	if err != nil {
		t.Fatalf("failed to list problems: %v", err)
	}

	if len(problems) < 4 {
		t.Errorf("expected at least 4 problems, got %d", len(problems))
	}

	twoSum, err := repo.GetProblem(ctx, "two_sum")
	if err != nil {
		t.Fatalf("failed to get two_sum: %v", err)
	}
	if twoSum.Entrypoint != "solve" || len(twoSum.Tests) != 3 {
		t.Errorf("unexpected two_sum problem content: %+v", twoSum)
	}
}
