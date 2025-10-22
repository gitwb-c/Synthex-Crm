package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

type PipelineRepository struct {
	client *ent.Client
}

func NewPipelineRepository(client *ent.Client) *PipelineRepository {
	return &PipelineRepository{
		client: client,
	}
}

func (s *PipelineRepository) Read(ctx context.Context) ([]*ent.Pipeline, error) {
	return s.client.Pipeline.Query().All(ctx)
}

func (s *PipelineRepository) Create(ctx context.Context, input ent.CreatePipelineInput) (*ent.Pipeline, error) {
	return s.client.Pipeline.Create().SetInput(input).Save(ctx)
}

func (s *PipelineRepository) UpdateID(ctx context.Context, id string, input ent.UpdatePipelineInput) (*ent.Pipeline, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Pipeline.UpdateOneID(uuidId).SetInput(input).Save(ctx)
}
func (s *PipelineRepository) Delete(ctx context.Context) {}
