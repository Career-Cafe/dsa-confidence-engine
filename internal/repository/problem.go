package repository

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
	"gopkg.in/yaml.v3"
)

func LoadProblemsFromDirectory(ctx context.Context, dir string, repo *SQLiteRepository) error {
	entries, err := os.ReadDir(dir)
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

		filePath := filepath.Join(dir, entry.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read problem file %s: %w", filePath, err)
		}

		var p model.Problem
		if err := yaml.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("failed to parse problem YAML %s: %w", filePath, err)
		}

		if p.ID == "" {
			p.ID = strings.TrimSuffix(entry.Name(), ext)
		}

		if err := repo.SaveProblem(ctx, &p); err != nil {
			return fmt.Errorf("failed to persist problem %s: %w", p.ID, err)
		}
	}

	return nil
}
