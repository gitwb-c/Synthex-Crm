package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/stage"
	"github.com/google/uuid"
)

type StageRepository struct {
	client *ent.Client
}

func NewStageRepository(client *ent.Client) *StageRepository {
	return &StageRepository{
		client: client,
	}
}

func (s *StageRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Stage, error) {
	return s.client.Stage.Query().Where(stage.TenantIdEQ(tenantId)).All(ctx)
}

func (s *StageRepository) Create(ctx context.Context, input ent.CreateStageInput) (*ent.Stage, error) {
	return s.client.Stage.Create().SetInput(input).Save(ctx)
}

func (s *StageRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateStageInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Stage.Update().Where(stage.IDIn(ids...)).SetInput(input).Save(ctx)
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

func (s *StageRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Stage.Delete().Where(stage.IDIn(ids...)).Exec(ctx)
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
