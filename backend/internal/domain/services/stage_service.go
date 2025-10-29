package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type StageService struct {
	repository *repositories.StageRepository
}

func NewStageService(repository *repositories.StageRepository) *StageService {
	return &StageService{
		repository: repository,
	}
}

func (s *StageService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Stage, error) {
	stage, err := s.repository.Read(ctx, tenantId)
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

func (s *StageService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateStageInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}
func (s *StageService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
