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

func (s *PipelineRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Pipeline, error) {
	return s.client.Pipeline.Query().Where(pipeline.TenantIdEQ(tenantId)).All(ctx)
}

func (s *PipelineRepository) Create(ctx context.Context, input ent.CreatePipelineInput) (*ent.Pipeline, error) {
	return s.client.Pipeline.Create().SetInput(input).Save(ctx)
}

func (s *PipelineRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdatePipelineInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Pipeline.Update().Where(pipeline.IDIn(ids...)).SetInput(input).Save(ctx)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	return n, nil
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
