package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type PipelineService struct {
	repository *repositories.PipelineRepository
}

func NewPipelineService(repository *repositories.PipelineRepository) *PipelineService {
	return &PipelineService{
		repository: repository,
	}
}

func (s *PipelineService) Read(ctx context.Context) ([]*ent.Pipeline, error) {
	pipeline, err := s.repository.Read(ctx)
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

func (s *PipelineService) UpdateID(ctx context.Context, id string, input ent.UpdatePipelineInput) (*ent.Pipeline, error) {
	pipeline, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return pipeline, nil
}
func (s *PipelineService) DeleteID(ctx context.Context, id string) error {
	err := s.repository.DeleteID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
