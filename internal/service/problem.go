package service

import (
	"context"

	"github.com/MishraShardendu22/dsa-confidence-engine/internal/model"
	"github.com/MishraShardendu22/dsa-confidence-engine/internal/repository"
)

type ProblemService struct {
	repo *repository.SQLiteRepository
}

func NewProblemService(repo *repository.SQLiteRepository) *ProblemService {
	return &ProblemService{repo: repo}
}

func (s *ProblemService) GetProblem(ctx context.Context, id string) (*model.Problem, error) {
	return s.repo.GetProblem(ctx, id)
}

func (s *ProblemService) ListProblems(ctx context.Context) ([]model.Problem, error) {
	return s.repo.ListProblems(ctx)
}

func (s *ProblemService) CreateProblem(ctx context.Context, p *model.Problem) error {
	return s.repo.SaveProblem(ctx, p)
}
