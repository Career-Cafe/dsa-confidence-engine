package nlp_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/dsa"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/nlp"
)

func TestCascadeMatcher(t *testing.T) {
	repoPath := filepath.Join("..", "..", "data", "dsa")
	repo, err := dsa.NewRepository(repoPath)
	if err != nil {
		t.Fatalf("failed to load ontology: %v", err)
	}

	embedder := nlp.NewLocalHashingEmbedder(256)
	ctx := context.Background()

	matcher, err := nlp.NewCascadeMatcher(ctx, repo.Ontology(), embedder, true)
	if err != nil {
		t.Fatalf("failed to initialize cascade matcher: %v", err)
	}

	// 1. Test Hashmap & Frequency Count explanation
	t.Run("hashmap and frequency counting", func(t *testing.T) {
		text := "I'll store each number's frequency count in a hashmap and then look up the values I need."
		claimed, err := matcher.Match(ctx, text)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var hasMap, hasFreq bool
		for _, c := range claimed {
			if c.ConceptID == "hashmap" {
				hasMap = true
			}
			if c.ConceptID == "frequency_count" {
				hasFreq = true
			}
		}

		if !hasMap {
			t.Errorf("expected hashmap to be claimed")
		}
		if !hasFreq {
			t.Errorf("expected frequency_count to be claimed")
		}
	})

	// 2. Test Sorting and Two Pointers opposite end
	t.Run("sorting and two pointers opposite end", func(t *testing.T) {
		text := "I'll sort the array and then use two pointers from opposite ends."
		claimed, err := matcher.Match(ctx, text)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var hasSort, hasTwoPtr, hasOpp bool
		for _, c := range claimed {
			if c.ConceptID == "sorting" {
				hasSort = true
			}
			if c.ConceptID == "two_pointers" {
				hasTwoPtr = true
			}
			if c.ConceptID == "opposite_end_pointers" {
				hasOpp = true
			}
		}

		if !hasSort {
			t.Errorf("expected sorting to be claimed")
		}
		if !hasTwoPtr {
			t.Errorf("expected two_pointers to be claimed")
		}
		if !hasOpp {
			t.Errorf("expected opposite_end_pointers to be claimed")
		}
	})

	// 3. Test Memoization and DP
	t.Run("memoization and top down dp", func(t *testing.T) {
		text := "I will use memoization to cache subproblems with top down dp."
		claimed, err := matcher.Match(ctx, text)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var hasMemo, hasDP bool
		for _, c := range claimed {
			if c.ConceptID == "memoization" {
				hasMemo = true
			}
			if c.ConceptID == "dynamic_programming" || c.ConceptID == "top_down_dp" {
				hasDP = true
			}
		}

		if !hasMemo {
			t.Errorf("expected memoization to be claimed")
		}
		if !hasDP {
			t.Errorf("expected dynamic_programming to be claimed")
		}
	})
}
