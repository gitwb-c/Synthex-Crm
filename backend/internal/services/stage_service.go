package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type StageService struct {
	repository *repositories.StageRepository
}

func NewStageService(repository *repositories.StageRepository) *StageService {
	return &StageService{
		repository: repository,
	}
}

func (s *StageService) Read(ctx context.Context) ([]*ent.Stage, error) {
	stage, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return stage, nil
}

func (s *StageService) Create(ctx context.Context, input ent.CreateStageInput) (*ent.Stage, error) {
	stage, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return stage, nil
}

func (s *StageService) UpdateID(ctx context.Context, id string, input ent.UpdateStageInput) (*ent.Stage, error) {
	stage, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return stage, nil
}
func (s *StageService) Delete() {}
