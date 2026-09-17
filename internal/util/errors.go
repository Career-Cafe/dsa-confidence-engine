package util

import "errors"

var (
	ErrProblemNotFound     = errors.New("problem not found")
	ErrEvaluationNotFound  = errors.New("evaluation not found")
	ErrInvalidSubmission   = errors.New("invalid submission")
	ErrTestExecutionFailed = errors.New("test execution failed")
	ErrAnalysisFailed      = errors.New("static code analysis failed")
	ErrOntologyNotFound    = errors.New("ontology concept not found")
)
