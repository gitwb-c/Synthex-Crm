package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type PipelineRepository struct {
	client *ent.Client
}

func NewPipelineRepository() *PipelineRepository {
	return &PipelineRepository{
		client: db.Client,
	}
}

// ([]*ent.Pipeline, error)
func (s *PipelineRepository) Read(data dto.ReadPipelineReq) {}

func (s *PipelineRepository) Create(data dto.CreatePipelineReq) {}

func (s *PipelineRepository) Update(data dto.UpdatePipelineReq) {}

func (s *PipelineRepository) Delete(data dto.DeletePipelineReq) {}
