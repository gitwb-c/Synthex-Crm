package repositories

import (
	"crm.saas/backend/internal/db"
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
func (s *PipelineRepository) Read() {}

func (s *PipelineRepository) Create() {}

func (s *PipelineRepository) Update() {}

func (s *PipelineRepository) Delete() {}
