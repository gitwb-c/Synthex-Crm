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

func (s *StageRepository) Read(ctx context.Context) ([]*ent.Stage, error) {
	return s.client.Stage.Query().All(ctx)
}

func (s *StageRepository) Create(ctx context.Context, input ent.CreateStageInput) (*ent.Stage, error) {
	return s.client.Stage.Create().SetInput(input).Save(ctx)
}

func (s *StageRepository) UpdateID(ctx context.Context, id string, input ent.UpdateStageInput) (*ent.Stage, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Stage.UpdateOneID(uuidId).SetInput(input).Save(ctx)
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
