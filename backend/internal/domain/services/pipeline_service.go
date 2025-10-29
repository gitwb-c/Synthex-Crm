package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type PipelineService struct {
	repository *repositories.PipelineRepository
}

func NewPipelineService(repository *repositories.PipelineRepository) *PipelineService {
	return &PipelineService{
		repository: repository,
	}
}

func (s *PipelineService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Pipeline, error) {
	pipeline, err := s.repository.Read(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	return pipeline, nil
}

func (s *PipelineService) Create(ctx context.Context, input ent.CreatePipelineInput) (*ent.Pipeline, error) {
	pipeline, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return pipeline, nil
}

func (s *PipelineService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdatePipelineInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}
func (s *PipelineService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
