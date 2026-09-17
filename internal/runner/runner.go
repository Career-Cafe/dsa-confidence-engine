package runner

import (
	"context"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
)

type TestCaseDetail struct {
	ID             string `json:"id"`
	Passed         bool   `json:"passed"`
	ExpectedOutput string `json:"expected"`
	ActualOutput   string `json:"actual"`
	Error          string `json:"error,omitempty"`
}

type TestResult struct {
	Status      model.TestStatus `json:"status"`
	PassedCount int              `json:"passed_count"`
	FailedCount int              `json:"failed_count"`
	TotalCount  int              `json:"total_count"`
	Details     []TestCaseDetail `json:"details"`
	Output      string           `json:"output,omitempty"`
	DurationMs  int64            `json:"duration_ms"`
	Error       string           `json:"error,omitempty"`
}

type Runner interface {
	Run(ctx context.Context, submission model.Submission, problem model.Problem) (TestResult, error)
}
