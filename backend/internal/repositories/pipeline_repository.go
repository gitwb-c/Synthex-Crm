package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/pipeline"
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
func (s *PipelineRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Pipeline.Delete().Where(pipeline.IDIn(ids...)).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
