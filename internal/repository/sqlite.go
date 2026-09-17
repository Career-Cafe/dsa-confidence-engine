package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/util"
	_ "modernc.org/sqlite"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(dbPath string) (*SQLiteRepository, error) {
	if dir := filepath.Dir(dbPath); dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create db directory: %w", err)
		}
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	repo := &SQLiteRepository{db: db}
	if err := repo.migrate(); err != nil {
		db.Close()
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return repo, nil
}

func (r *SQLiteRepository) Close() error {
	return r.db.Close()
}

func (r *SQLiteRepository) migrate() error {
	schema := `
	CREATE TABLE IF NOT EXISTS problems (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		language TEXT NOT NULL,
		entrypoint TEXT NOT NULL,
		entrypoint_aliases_json TEXT DEFAULT '[]',
		starter_code TEXT DEFAULT '',
		tests_json TEXT NOT NULL,
		accepted_strategies_json TEXT NOT NULL,
		required_concepts_json TEXT NOT NULL,
		optional_concepts_json TEXT NOT NULL,
		primary_concepts_json TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS evaluations (
		id TEXT PRIMARY KEY,
		problem_id TEXT NOT NULL,
		submission_code TEXT NOT NULL,
		explanation TEXT NOT NULL,
		test_result TEXT NOT NULL,
		passed_tests INTEGER NOT NULL,
		failed_tests INTEGER NOT NULL,
		total_tests INTEGER NOT NULL,
		actual_concepts_json TEXT NOT NULL,
		claimed_concepts_json TEXT NOT NULL,
		matched_concepts_json TEXT NOT NULL,
		missing_concepts_json TEXT NOT NULL,
		extra_concepts_json TEXT NOT NULL,
		fidelity_score REAL NOT NULL,
		decision TEXT NOT NULL,
		reason TEXT,
		evidence_json TEXT NOT NULL,
		diagnostics_json TEXT NOT NULL,
		duration_ms INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := r.db.Exec(schema); err != nil {
		return err
	}
	// Migrations for existing databases
	_, _ = r.db.Exec("ALTER TABLE problems ADD COLUMN entrypoint_aliases_json TEXT DEFAULT '[]';")
	_, _ = r.db.Exec("ALTER TABLE problems ADD COLUMN starter_code TEXT DEFAULT '';")
	return nil
}

func (r *SQLiteRepository) SaveProblem(ctx context.Context, p *model.Problem) error {
	aliasesJSON, err := json.Marshal(p.EntrypointAliases)
	if err != nil {
		return err
	}
	testsJSON, err := json.Marshal(p.Tests)
	if err != nil {
		return err
	}
	strategiesJSON, err := json.Marshal(p.AcceptedStrategies)
	if err != nil {
		return err
	}
	reqJSON, err := json.Marshal(p.RequiredConcepts)
	if err != nil {
		return err
	}
	optJSON, err := json.Marshal(p.OptionalConcepts)
	if err != nil {
		return err
	}
	priJSON, err := json.Marshal(p.PrimaryConcepts)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO problems (
		id, title, description, language, entrypoint,
		entrypoint_aliases_json, starter_code,
		tests_json, accepted_strategies_json, required_concepts_json,
		optional_concepts_json, primary_concepts_json
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(id) DO UPDATE SET
		title=excluded.title,
		description=excluded.description,
		language=excluded.language,
		entrypoint=excluded.entrypoint,
		entrypoint_aliases_json=excluded.entrypoint_aliases_json,
		starter_code=excluded.starter_code,
		tests_json=excluded.tests_json,
		accepted_strategies_json=excluded.accepted_strategies_json,
		required_concepts_json=excluded.required_concepts_json,
		optional_concepts_json=excluded.optional_concepts_json,
		primary_concepts_json=excluded.primary_concepts_json;
	`
	_, err = r.db.ExecContext(ctx, query,
		p.ID, p.Title, p.Description, p.Language, p.Entrypoint,
		string(aliasesJSON), p.StarterCode,
		string(testsJSON), string(strategiesJSON), string(reqJSON),
		string(optJSON), string(priJSON),
	)
	return err
}

func (r *SQLiteRepository) GetProblem(ctx context.Context, id string) (*model.Problem, error) {
	query := `
	SELECT id, title, description, language, entrypoint,
	       entrypoint_aliases_json, starter_code,
	       tests_json, accepted_strategies_json, required_concepts_json,
	       optional_concepts_json, primary_concepts_json
	FROM problems WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var p model.Problem
	var aliasesJSON, testsJSON, stratJSON, reqJSON, optJSON, priJSON string
	err := row.Scan(
		&p.ID, &p.Title, &p.Description, &p.Language, &p.Entrypoint,
		&aliasesJSON, &p.StarterCode,
		&testsJSON, &stratJSON, &reqJSON, &optJSON, &priJSON,
	)
	if err == sql.ErrNoRows {
		return nil, util.ErrProblemNotFound
	}
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal([]byte(aliasesJSON), &p.EntrypointAliases)
	if err := json.Unmarshal([]byte(testsJSON), &p.Tests); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(stratJSON), &p.AcceptedStrategies); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(reqJSON), &p.RequiredConcepts); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(optJSON), &p.OptionalConcepts); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(priJSON), &p.PrimaryConcepts); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *SQLiteRepository) ListProblems(ctx context.Context) ([]model.Problem, error) {
	query := `
	SELECT id, title, description, language, entrypoint,
	       entrypoint_aliases_json, starter_code,
	       tests_json, accepted_strategies_json, required_concepts_json,
	       optional_concepts_json, primary_concepts_json
	FROM problems ORDER BY id ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []model.Problem
	for rows.Next() {
		var p model.Problem
		var aliasesJSON, testsJSON, stratJSON, reqJSON, optJSON, priJSON string
		if err := rows.Scan(
			&p.ID, &p.Title, &p.Description, &p.Language, &p.Entrypoint,
			&aliasesJSON, &p.StarterCode,
			&testsJSON, &stratJSON, &reqJSON, &optJSON, &priJSON,
		); err != nil {
			return nil, err
		}
		_ = json.Unmarshal([]byte(aliasesJSON), &p.EntrypointAliases)
		_ = json.Unmarshal([]byte(testsJSON), &p.Tests)
		_ = json.Unmarshal([]byte(stratJSON), &p.AcceptedStrategies)
		_ = json.Unmarshal([]byte(reqJSON), &p.RequiredConcepts)
		_ = json.Unmarshal([]byte(optJSON), &p.OptionalConcepts)
		_ = json.Unmarshal([]byte(priJSON), &p.PrimaryConcepts)
		problems = append(problems, p)
	}
	return problems, rows.Err()
}

func (r *SQLiteRepository) SaveEvaluation(ctx context.Context, e *model.Evaluation) error {
	actualJSON, _ := json.Marshal(e.ActualConcepts)
	claimedJSON, _ := json.Marshal(e.ClaimedConcepts)
	matchedJSON, _ := json.Marshal(e.MatchedConcepts)
	missingJSON, _ := json.Marshal(e.MissingConcepts)
	extraJSON, _ := json.Marshal(e.ExtraConcepts)
	evidenceJSON, _ := json.Marshal(e.Evidence)
	diagJSON, _ := json.Marshal(e.Diagnostics)

	query := `
	INSERT INTO evaluations (
		id, problem_id, submission_code, explanation, test_result,
		passed_tests, failed_tests, total_tests,
		actual_concepts_json, claimed_concepts_json, matched_concepts_json,
		missing_concepts_json, extra_concepts_json, fidelity_score,
		decision, reason, evidence_json, diagnostics_json,
		duration_ms, created_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	createdAt := e.CreatedAt
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	_, err := r.db.ExecContext(ctx, query,
		e.ID, e.ProblemID, e.SubmissionCode, e.Explanation, string(e.TestResult),
		e.PassedTests, e.FailedTests, e.TotalTests,
		string(actualJSON), string(claimedJSON), string(matchedJSON),
		string(missingJSON), string(extraJSON), e.FidelityScore,
		string(e.Decision), e.Reason, string(evidenceJSON), string(diagJSON),
		e.DurationMs, createdAt,
	)
	return err
}

func (r *SQLiteRepository) GetEvaluation(ctx context.Context, id string) (*model.Evaluation, error) {
	query := `
	SELECT id, problem_id, submission_code, explanation, test_result,
	       passed_tests, failed_tests, total_tests,
	       actual_concepts_json, claimed_concepts_json, matched_concepts_json,
	       missing_concepts_json, extra_concepts_json, fidelity_score,
	       decision, reason, evidence_json, diagnostics_json,
	       duration_ms, created_at
	FROM evaluations WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var e model.Evaluation
	var testRes, decision string
	var actualJSON, claimedJSON, matchedJSON, missingJSON, extraJSON, evidenceJSON, diagJSON string
	var reason sql.NullString

	err := row.Scan(
		&e.ID, &e.ProblemID, &e.SubmissionCode, &e.Explanation, &testRes,
		&e.PassedTests, &e.FailedTests, &e.TotalTests,
		&actualJSON, &claimedJSON, &matchedJSON,
		&missingJSON, &extraJSON, &e.FidelityScore,
		&decision, &reason, &evidenceJSON, &diagJSON,
		&e.DurationMs, &e.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, util.ErrEvaluationNotFound
	}
	if err != nil {
		return nil, err
	}

	e.TestResult = model.TestStatus(testRes)
	e.Decision = model.Decision(decision)
	if reason.Valid {
		e.Reason = reason.String
	}

	_ = json.Unmarshal([]byte(actualJSON), &e.ActualConcepts)
	_ = json.Unmarshal([]byte(claimedJSON), &e.ClaimedConcepts)
	_ = json.Unmarshal([]byte(matchedJSON), &e.MatchedConcepts)
	_ = json.Unmarshal([]byte(missingJSON), &e.MissingConcepts)
	_ = json.Unmarshal([]byte(extraJSON), &e.ExtraConcepts)
	_ = json.Unmarshal([]byte(evidenceJSON), &e.Evidence)
	_ = json.Unmarshal([]byte(diagJSON), &e.Diagnostics)

	return &e, nil
}
