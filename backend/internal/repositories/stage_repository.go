package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
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

func (s *StageRepository) DeleteID(ctx context.Context, id string) error {
	uuidId, e := uuid.Parse(id)
	if e != nil {
		return e
	}
	err := s.client.Stage.DeleteOneID(uuidId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
