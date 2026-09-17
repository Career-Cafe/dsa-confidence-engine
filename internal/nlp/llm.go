package nlp

import (
	"context"
	"errors"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
)

var ErrLLMDisabled = errors.New("LLM escalation is disabled")

type LLMPayload struct {
	CandidateExplanation string   `json:"candidate_explanation"`
	ClaimedConcepts      []string `json:"claimed_concepts"`
	DetectedConcepts     []string `json:"detected_concepts"`
	Evidence             []string `json:"evidence"`
}

type LLMDecision struct {
	ResolvedClaimedConcepts []string `json:"resolved_claimed_concepts"`
	Confidence              float64  `json:"confidence"`
	Reasoning               string   `json:"reasoning"`
}

type LLMEscalator interface {
	Escalate(ctx context.Context, payload LLMPayload) (*LLMDecision, error)
}

type DisabledLLMEscalator struct{}

func (d *DisabledLLMEscalator) Escalate(ctx context.Context, payload LLMPayload) (*LLMDecision, error) {
	return nil, ErrLLMDisabled
}

func FormatEvidenceSummaries(evidence []model.Evidence) []string {
	var summaries []string
	for _, e := range evidence {
		if e.Reachable {
			summaries = append(summaries, e.Description)
		}
	}
	return summaries
}
