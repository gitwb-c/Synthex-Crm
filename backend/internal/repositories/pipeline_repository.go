package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type PipelineRepository struct {
	client *ent.Client
}

func NewPipelineRepository(client *ent.Client) *PipelineRepository {
	return &PipelineRepository{
		client: client,
	}
}

// ([]*ent.Pipeline, error)
func (s *PipelineRepository) Read(ctx context.Context) {}

func (s *PipelineRepository) Create(ctx context.Context) {}

func (s *PipelineRepository) Update(ctx context.Context) {}

func (s *PipelineRepository) Delete(ctx context.Context) {}
