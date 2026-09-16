package dsa_test

import (
	"path/filepath"
	"testing"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/dsa"
)

func TestOntologyRepositoryLoading(t *testing.T) {
	repoPath := filepath.Join("..", "..", "data", "dsa")
	repo, err := dsa.NewRepository(repoPath)
	if err != nil {
		t.Fatalf("failed to load ontology repository: %v", err)
	}

	ont := repo.Ontology()
	concepts := ont.AllConcepts()
	if len(concepts) < 29 {
		t.Fatalf("expected at least 29 concepts, got %d", len(concepts))
	}

	// 1. Test Hashmap resolution across aliases and synonyms
	hashmapAliases := []string{
		"hashmap",
		"hash map",
		"dictionary",
		"lookup table",
		"frequency map",
		"frequency table",
	}

	for _, alias := range hashmapAliases {
		c, ok := ont.FindConceptByAlias(alias)
		if !ok {
			t.Errorf("expected alias %q to resolve to a concept", alias)
			continue
		}
		if c.ID != "hashmap" && c.Parent != "hashing" && c.Parent != "hashmap" {
			t.Errorf("expected alias %q to resolve to hashmap or child, got %s (parent: %s)", alias, c.ID, c.Parent)
		}
	}

	// 2. Test 1-D DP resolution
	dp1DAliases := []string{
		"1-d dp",
		"1d dp",
		"one dimensional dynamic programming",
		"linear dp",
	}

	for _, alias := range dp1DAliases {
		c, ok := ont.FindConceptByAlias(alias)
		if !ok {
			t.Errorf("expected DP alias %q to resolve", alias)
			continue
		}
		if c.ID != "dp_1d" {
			t.Errorf("expected alias %q to resolve to dp_1d, got %s", alias, c.ID)
		}
	}

	// 3. Test Ancestor and Lineage checks
	if !ont.IsAncestor("dynamic_programming", "dp_1d") {
		t.Errorf("expected dynamic_programming to be ancestor of dp_1d")
	}
	if !ont.IsAncestor("hashing", "frequency_count") {
		t.Errorf("expected hashing to be ancestor of frequency_count")
	}

	lineage := ont.GetLineage("dp_string_edit_distance")
	if len(lineage) < 3 {
		t.Errorf("expected edit distance lineage depth >= 3, got %v", lineage)
	}
}
