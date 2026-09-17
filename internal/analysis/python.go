package analysis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
)

type PythonAnalyzer struct {
	scriptPath string
	registry   *DetectorRegistry
}

func NewPythonAnalyzer(scriptPath string) *PythonAnalyzer {
	if scriptPath == "" {
		// Default to relative internal/analysis/python_ast.py
		scriptPath = filepath.Join("internal", "analysis", "python_ast.py")
	}
	return &PythonAnalyzer{
		scriptPath: scriptPath,
		registry:   NewDetectorRegistry(),
	}
}

type pythonASTResponse struct {
	Status    string         `json:"status"`
	Functions []FunctionNode `json:"functions"`
	Error     string         `json:"error,omitempty"`
}

func (a *PythonAnalyzer) Analyze(ctx context.Context, source []byte, contract AnalysisContract) (AnalysisResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Locate python_ast.py if running from different working directories
	script := a.scriptPath
	if _, err := os.Stat(script); os.IsNotExist(err) {
		// Try finding it relative to project root
		candidate := filepath.Join("..", "..", "internal", "analysis", "python_ast.py")
		if _, err := os.Stat(candidate); err == nil {
			script = candidate
		}
	}

	cmd := exec.CommandContext(ctx, "python3", script)
	cmd.Stdin = bytes.NewReader(source)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return AnalysisResult{
			Status: "UNKNOWN",
			Error:  fmt.Sprintf("AST extraction failed: %v (stderr: %s)", err, stderr.String()),
		}, nil
	}

	var resp pythonASTResponse
	if err := json.Unmarshal(stdout.Bytes(), &resp); err != nil {
		return AnalysisResult{
			Status: "UNKNOWN",
			Error:  fmt.Sprintf("Failed to parse AST response: %v", err),
		}, nil
	}

	if resp.Status == "ERROR" {
		return AnalysisResult{
			Status: "UNKNOWN",
			Error:  resp.Error,
		}, nil
	}

	// 1. Call graph & Reachability
	reachableFuncs := ComputeReachableFunctions(resp.Functions, contract)

	// 2. Backward Output Relevance
	relevanceMap := ComputeOutputRelevance(reachableFuncs)

	// 3. Concept Detectors
	rawConcepts := a.registry.RunAll(resp.Functions, reachableFuncs, relevanceMap)

	// 4. Deduplicate and aggregate evidence by ConceptID
	conceptMap := make(map[string]*model.DetectedConcept)
	var allEvidence []model.Evidence

	for _, c := range rawConcepts {
		// Only keep concept if it has evidence
		if len(c.Evidence) == 0 {
			continue
		}

		if existing, ok := conceptMap[c.ConceptID]; ok {
			existing.Evidence = append(existing.Evidence, c.Evidence...)
			if c.Reachable {
				existing.Reachable = true
			}
			if c.OutputRelevant {
				existing.OutputRelevant = true
			}
		} else {
			copyC := c
			conceptMap[c.ConceptID] = &copyC
		}
		allEvidence = append(allEvidence, c.Evidence...)
	}

	detectedList := make([]model.DetectedConcept, 0, len(conceptMap))
	for _, c := range conceptMap {
		detectedList = append(detectedList, *c)
	}

	return AnalysisResult{
		ActualConcepts: detectedList,
		Evidence:       allEvidence,
		Status:         "SUCCESS",
	}, nil
}
