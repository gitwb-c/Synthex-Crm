package services

import (
	"crm.saas/backend/internal/dto"
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
func (s *PipelineService) Read(data dto.ReadPipelineReq) {}

func (s *PipelineService) Create(data dto.CreatePipelineReq) {}

func (s *PipelineService) Update(data dto.UpdatePipelineReq) {}

func (s *PipelineService) Delete(data dto.DeletePipelineReq) {}
