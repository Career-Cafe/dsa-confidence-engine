package fidelity_test

import (
	"path/filepath"
	"testing"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/dsa"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/fidelity"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
)

func TestScorer(t *testing.T) {
	repoPath := filepath.Join("..", "..", "data", "dsa")
	repo, err := dsa.NewRepository(repoPath)
	if err != nil {
		t.Fatalf("failed to load ontology: %v", err)
	}

	thresholds := fidelity.DefaultThresholds()
	scorer := fidelity.NewScorer(repo.Ontology(), thresholds)

	problem := model.Problem{
		ID:               "two_sum",
		PrimaryConcepts:  []string{"hashmap"},
		RequiredConcepts: []string{"hashmap"},
		OptionalConcepts: []string{"two_pointers"},
	}

	// 1. Strong Match -> ACCEPT
	t.Run("strong match output relevant -> ACCEPT", func(t *testing.T) {
		claimed := []model.ClaimedConcept{
			{ConceptID: "hashmap", Name: "Hash Map"},
		}
		actual := []model.DetectedConcept{
			{
				ConceptID:      "hashmap",
				Name:           "Hash Map",
				Reachable:      true,
				OutputRelevant: true,
			},
		}

		res := scorer.Score(actual, claimed, problem)
		if res.Decision != model.DecisionAccept {
			t.Errorf("expected ACCEPT, got %s (score: %f)", res.Decision, res.Score)
		}
		if res.Score < 0.95 {
			t.Errorf("expected score >= 0.95, got %f", res.Score)
		}
	})

	// 2. Claimed two concepts, one output relevant, one not output relevant -> REJUSTIFY
	t.Run("two_pointers + non-output-relevant hashmap -> REJUSTIFY", func(t *testing.T) {
		claimed := []model.ClaimedConcept{
			{ConceptID: "two_pointers", Name: "Two Pointers"},
			{ConceptID: "hashmap", Name: "Hash Map"},
		}
		actual := []model.DetectedConcept{
			{
				ConceptID:      "two_pointers",
				Name:           "Two Pointers",
				Reachable:      true,
				OutputRelevant: true,
			},
			{
				ConceptID:      "hashmap",
				Name:           "Hash Map",
				Reachable:      true,
				OutputRelevant: false, // NOT output relevant
			},
		}

		res := scorer.Score(actual, claimed, problem)
		if res.Decision != model.DecisionRejustify {
			t.Errorf("expected REJUSTIFY, got %s (score: %f)", res.Decision, res.Score)
		}
		if res.Score < 0.90 || res.Score >= 0.95 {
			t.Errorf("expected score in [0.90, 0.95), got %f", res.Score)
		}
	})

	// 3. Completely Missing Primary Concept -> REJECT
	t.Run("missing primary concept -> REJECT", func(t *testing.T) {
		claimed := []model.ClaimedConcept{
			{ConceptID: "hashmap", Name: "Hash Map"},
		}
		actual := []model.DetectedConcept{
			{
				ConceptID:      "sorting",
				Name:           "Sorting",
				Reachable:      true,
				OutputRelevant: true,
			},
		}

		res := scorer.Score(actual, claimed, problem)
		if res.Decision != model.DecisionReject {
			t.Errorf("expected REJECT, got %s (score: %f)", res.Decision, res.Score)
		}
		if res.Score >= 0.90 {
			t.Errorf("expected score < 0.90, got %f", res.Score)
		}
	})
}
