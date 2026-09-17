package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/analysis"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/api"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/config"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/controller"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/dsa"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/evaluator"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/fidelity"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/nlp"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/repository"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/runner"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("[FATAL] failed to load configuration: %v", err)
	}

	ctx := context.Background()

	// 1. Initialize SQLite repository
	sqliteRepo, err := repository.NewSQLiteRepository(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("[FATAL] failed to initialize sqlite repository: %v", err)
	}
	defer sqliteRepo.Close()

	// 2. Load problem definitions into SQLite
	if err := repository.LoadProblemsFromDirectory(ctx, cfg.ProblemsPath, sqliteRepo); err != nil {
		log.Printf("[WARN] failed to load problems from %s: %v", cfg.ProblemsPath, err)
	}

	// 3. Initialize DSA Ontology
	ontologyRepo, err := dsa.NewRepository(cfg.OntologyPath)
	if err != nil {
		log.Fatalf("[FATAL] failed to load DSA ontology: %v", err)
	}
	log.Printf("[INFO] Loaded %d DSA ontology concepts across %s", len(ontologyRepo.Ontology().AllConcepts()), cfg.OntologyPath)

	// 4. Initialize Local Embedder and Cascade Matcher
	var embedder nlp.Embedder
	if cfg.EmbeddingEnabled {
		embedder = nlp.NewLocalHashingEmbedder(256)
	}
	matcher, err := nlp.NewCascadeMatcher(ctx, ontologyRepo.Ontology(), embedder, cfg.EmbeddingEnabled)
	if err != nil {
		log.Fatalf("[FATAL] failed to initialize cascade matcher: %v", err)
	}

	// 5. Initialize Subprocess Test Runner
	testRunner := runner.NewLocalRunner(cfg.RunnerTimeoutMs)

	// 6. Initialize Point A Python Analyzer
	pythonAnalyzer := analysis.NewPythonAnalyzer("")

	// 7. Initialize Fidelity Scorer
	thresholds := fidelity.Thresholds{
		AcceptThreshold:    cfg.AcceptThreshold,
		RejustifyThreshold: cfg.RejustifyThreshold,
	}
	scorer := fidelity.NewScorer(ontologyRepo.Ontology(), thresholds)

	// 8. Initialize Core Evaluator Engine
	evalEngine := evaluator.NewEvaluator(testRunner, pythonAnalyzer, matcher, scorer, nil)

	// 9. Services & Controllers
	probService := service.NewProblemService(sqliteRepo)
	evalService := service.NewEvaluationService(evalEngine, sqliteRepo)

	probCtrl := controller.NewProblemController(probService)
	evalCtrl := controller.NewEvaluationController(evalService)

	// 10. Fiber Application
	app := fiber.New(fiber.Config{
		AppName:      "DSA Approach Fidelity Evaluator",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// Register Routes
	api.RegisterRoutes(app, api.RouterConfig{
		ProblemHandler:    probCtrl,
		EvaluationHandler: evalCtrl,
		ProblemService:    probService,
		EvaluationService: evalService,
		OntologyRepo:      ontologyRepo,
	})

	// Graceful shutdown channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("[INFO] Shutting down HTTP server...")
		_ = app.Shutdown()
	}()

	log.Printf("[INFO] Server starting on %s (env: %s)", cfg.AppAddr, cfg.AppEnv)
	if err := app.Listen(cfg.AppAddr); err != nil {
		fmt.Printf("[INFO] Server shut down: %v\n", err)
	}
}
