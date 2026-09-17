package dsa

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Repository struct {
	ontologyPath string
	ontology     *Ontology
}

func NewRepository(ontologyPath string) (*Repository, error) {
	repo := &Repository{
		ontologyPath: ontologyPath,
		ontology:     NewOntology(),
	}

	if err := repo.load(); err != nil {
		return nil, fmt.Errorf("failed to load ontology: %w", err)
	}

	return repo, nil
}

func (r *Repository) Ontology() *Ontology {
	return r.ontology
}

type yamlConceptWrapper struct {
	Concepts []Concept `yaml:"concepts"`
}

func (r *Repository) load() error {
	entries, err := os.ReadDir(r.ontologyPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".yaml" && ext != ".yml" {
			continue
		}

		filePath := filepath.Join(r.ontologyPath, entry.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", filePath, err)
		}

		var concepts []Concept
		if err := yaml.Unmarshal(data, &concepts); err != nil {
			var wrapper yamlConceptWrapper
			if err2 := yaml.Unmarshal(data, &wrapper); err2 != nil {
				return fmt.Errorf("failed to parse yaml in %s: %w", filePath, err)
			}
			concepts = wrapper.Concepts
		}

		categoryName := strings.TrimSuffix(entry.Name(), ext)
		for i := range concepts {
			c := concepts[i]
			if c.Category == "" {
				c.Category = categoryName
			}
			r.ontology.AddConcept(&c)
		}
	}

	return nil
}
