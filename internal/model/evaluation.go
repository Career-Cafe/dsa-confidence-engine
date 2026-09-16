package model

import "time"

type Decision string

const (
	DecisionAccept    Decision = "ACCEPT"
	DecisionRejustify Decision = "REJUSTIFY"
	DecisionReject    Decision = "REJECT"
)

type TestStatus string

const (
	TestStatusPass TestStatus = "PASS"
	TestStatusFail TestStatus = "FAIL"
)

type Evidence struct {
	Type           string `json:"type"`
	File           string `json:"file,omitempty"`
	Line           int    `json:"line"`
	Description    string `json:"description"`
	Reachable      bool   `json:"reachable"`
	OutputRelevant bool   `json:"output_relevant"`
}

type ConceptRole string

const (
	RolePrimary    ConceptRole = "PRIMARY"
	RoleSupporting ConceptRole = "SUPPORTING"
	RoleAuxiliary  ConceptRole = "AUXILIARY"
	RoleIncidental ConceptRole = "INCIDENTAL"
)

type DetectedConcept struct {
	ConceptID      string      `json:"id"`
	Name           string      `json:"name,omitempty"`
	Category       string      `json:"category,omitempty"`
	Evidence       []Evidence  `json:"evidence,omitempty"`
	Reachable      bool        `json:"reachable"`
	OutputRelevant bool        `json:"output_relevant"`
	Confidence     float64     `json:"confidence"`
	Role           ConceptRole `json:"role,omitempty"`
	Status         string      `json:"status,omitempty"`
}

type ClaimedConcept struct {
	ConceptID     string      `json:"id"`
	Name          string      `json:"name,omitempty"`
	Category      string      `json:"category,omitempty"`
	Confidence    float64     `json:"confidence"`
	MatchedPhrase string      `json:"matched_phrase,omitempty"`
	Stage         string      `json:"stage,omitempty"`
	Role          ConceptRole `json:"role,omitempty"`
}

type ConceptMatch struct {
	ConceptID      string      `json:"concept_id"`
	Name           string      `json:"name,omitempty"`
	Role           ConceptRole `json:"role,omitempty"`
	OutputRelevant bool        `json:"output_relevant"`
	Contribution   float64     `json:"contribution"`
	Notes          string      `json:"notes,omitempty"`
}

type Evaluation struct {
	ID              string            `json:"id"`
	ProblemID       string            `json:"problem_id"`
	SubmissionCode  string            `json:"submission_code,omitempty"`
	Explanation     string            `json:"explanation,omitempty"`
	TestResult      TestStatus        `json:"test_result"`
	PassedTests     int               `json:"passed_tests"`
	FailedTests     int               `json:"failed_tests"`
	TotalTests      int               `json:"total_tests"`
	ActualConcepts  []DetectedConcept `json:"actual_concepts"`
	ClaimedConcepts []ClaimedConcept  `json:"claimed_concepts"`
	MatchedConcepts []ConceptMatch    `json:"matched_concepts"`
	MissingConcepts []ConceptMatch    `json:"missing_concepts"`
	ExtraConcepts   []ConceptMatch    `json:"extra_concepts"`
	FidelityScore   float64           `json:"fidelity_score"`
	Decision        Decision          `json:"decision"`
	Reason          string            `json:"reason,omitempty"`
	Evidence        []Evidence        `json:"evidence"`
	Diagnostics     []string          `json:"diagnostics"`
	DurationMs      int64             `json:"duration_ms"`
	CreatedAt       time.Time         `json:"created_at"`
}
