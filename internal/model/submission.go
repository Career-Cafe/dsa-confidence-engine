package model

type Submission struct {
	ProblemID   string `json:"problem_id"`
	SourceCode  string `json:"source_code"`
	Explanation string `json:"explanation"`
}
