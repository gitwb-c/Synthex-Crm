package services

import (
	"crm.saas/backend/internal/repositories"
)

type PipelineService struct {
	repository *repositories.PipelineRepository
}

func NewPipelineService(repository *repositories.PipelineRepository) *PipelineService {
	return &PipelineService{
		repository: repository,
	}
}

// ([]*ent.Pipeline, error)
func (s *PipelineService) Read() {}

func (s *PipelineService) Create() {}

func (s *PipelineService) Update() {}

func (s *PipelineService) Delete() {}
