package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/analysis"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/api"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/controller"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/dsa"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/evaluator"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/fidelity"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/nlp"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/repository"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/runner"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/service"
)

func setupTestApp(t *testing.T) (*fiber.App, *repository.SQLiteRepository) {
	tmpDir, err := os.MkdirTemp("", "ctrl-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(tmpDir) })

	dbPath := filepath.Join(tmpDir, "test.db")
	sqliteRepo, err := repository.NewSQLiteRepository(dbPath)
	if err != nil {
		t.Fatalf("failed to create sqlite repo: %v", err)
	}
	t.Cleanup(func() { sqliteRepo.Close() })

	ontologyPath := filepath.Join("..", "..", "data", "dsa")
	ontRepo, err := dsa.NewRepository(ontologyPath)
	if err != nil {
		t.Fatalf("failed to load ontology: %v", err)
	}

	problemsPath := filepath.Join("..", "..", "data", "problems")
	ctx := context.Background()
	if err := repository.LoadProblemsFromDirectory(ctx, problemsPath, sqliteRepo); err != nil {
		t.Fatalf("failed to load problems: %v", err)
	}

	embedder := nlp.NewLocalHashingEmbedder(256)
	matcher, err := nlp.NewCascadeMatcher(ctx, ontRepo.Ontology(), embedder, true)
	if err != nil {
		t.Fatalf("failed to create matcher: %v", err)
	}

	r := runner.NewLocalRunner(2000)
	scriptPath := filepath.Join("..", "analysis", "python_ast.py")
	a := analysis.NewPythonAnalyzer(scriptPath)
	s := fidelity.NewScorer(ontRepo.Ontology(), fidelity.DefaultThresholds())
	evalEngine := evaluator.NewEvaluator(r, a, matcher, s, nil)

	probService := service.NewProblemService(sqliteRepo)
	evalService := service.NewEvaluationService(evalEngine, sqliteRepo)

	probCtrl := controller.NewProblemController(probService)
	evalCtrl := controller.NewEvaluationController(evalService)

	app := fiber.New()
	api.RegisterRoutes(app, api.RouterConfig{
		ProblemHandler:    probCtrl,
		EvaluationHandler: evalCtrl,
		ProblemService:    probService,
		EvaluationService: evalService,
		OntologyRepo:      ontRepo,
	})

	return app, sqliteRepo
}

func TestHealthEndpoint(t *testing.T) {
	app, _ := setupTestApp(t)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	var envelope api.Response
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !envelope.Success {
		t.Errorf("expected success=true, got %v", envelope.Success)
	}
}

func TestProblemEndpoints(t *testing.T) {
	app, _ := setupTestApp(t)

	// 1. List Problems
	req := httptest.NewRequest(http.MethodGet, "/api/problems", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	// 2. Get Problem by ID
	req = httptest.NewRequest(http.MethodGet, "/api/problems/two_sum", nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestEvaluateEndpoint(t *testing.T) {
	app, _ := setupTestApp(t)

	submission := model.Submission{
		ProblemID: "two_sum",
		SourceCode: `
def solve(nums, target):
    seen = {}
    for i, n in enumerate(nums):
        diff = target - n
        if diff in seen:
            return [seen[diff], i]
        seen[n] = i
    return []
`,
		Explanation: "I'll store elements in a hashmap and check for complements in constant time.",
	}

	bodyBytes, _ := json.Marshal(submission)
	req := httptest.NewRequest(http.MethodPost, "/api/evaluate", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("evaluate request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var envelope struct {
		Success bool             `json:"success"`
		Data    model.Evaluation `json:"data"`
		Error   *api.APIError    `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		t.Fatalf("failed to decode eval response: %v", err)
	}

	if !envelope.Success {
		t.Fatalf("expected success=true, got error: %+v", envelope.Error)
	}
	if envelope.Data.Decision != model.DecisionAccept {
		t.Errorf("expected ACCEPT, got %s", envelope.Data.Decision)
	}
	if envelope.Data.TestResult != model.TestStatusPass {
		t.Errorf("expected PASS, got %s", envelope.Data.TestResult)
	}
	if envelope.Data.FidelityScore < 0.95 {
		t.Errorf("expected score >= 0.95, got %f", envelope.Data.FidelityScore)
	}

	// Fetch via GET /api/evaluations/:id
	req = httptest.NewRequest(http.MethodGet, "/api/evaluations/"+envelope.Data.ID, nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("get evaluation failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}
