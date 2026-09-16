package analysis

import (
	"context"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
)

type AnalysisContract struct {
	Entrypoint string `json:"entrypoint"`
	Language   string `json:"language"`
}

type AnalysisResult struct {
	ActualConcepts []model.DetectedConcept `json:"actual_concepts"`
	Evidence       []model.Evidence        `json:"evidence"`
	Status         string                  `json:"status"` // "SUCCESS", "UNKNOWN", "ERROR"
	Error          string                  `json:"error,omitempty"`
}

type Analyzer interface {
	Analyze(ctx context.Context, source []byte, contract AnalysisContract) (AnalysisResult, error)
}
