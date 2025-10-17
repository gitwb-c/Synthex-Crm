package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type PipelineService struct {
	client *ent.Client
}

func NewPipelineService() *PipelineService {
	return &PipelineService{
		client: db.Client,
	}
}

// ([]*ent.Pipeline, error)
func (s *PipelineService) Read(data dto.ReadPipelineReq) {}

func (s *PipelineService) Create(data dto.ReadPipelineReq) {}

func (s *PipelineService) Update(data dto.ReadPipelineReq) {}

func (s *PipelineService) Delete(data dto.ReadPipelineReq) {}
