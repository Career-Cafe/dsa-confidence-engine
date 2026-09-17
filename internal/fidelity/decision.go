package fidelity

import "github.com/MishraShardendu22/dsa-confidence-engine/internal/model"

type Thresholds struct {
	AcceptThreshold    float64
	RejustifyThreshold float64
}

func DefaultThresholds() Thresholds {
	return Thresholds{
		AcceptThreshold:    0.95,
		RejustifyThreshold: 0.90,
	}
}

func DetermineDecision(score float64, t Thresholds) model.Decision {
	if score >= t.AcceptThreshold {
		return model.DecisionAccept
	}
	if score >= t.RejustifyThreshold {
		return model.DecisionRejustify
	}
	return model.DecisionReject
}
