package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

type CostumerRepository struct {
	client *ent.Client
}

func NewCostumerRepository(client *ent.Client) *CostumerRepository {
	return &CostumerRepository{
		client: client,
	}
}

func (s *CostumerRepository) Read(ctx context.Context) ([]*ent.Costumer, error) {
	return s.client.Costumer.Query().All(ctx)
}

func (s *CostumerRepository) Create(ctx context.Context, input ent.CreateCostumerInput) (*ent.Costumer, error) {
	return s.client.Costumer.Create().SetInput(input).Save(ctx)
}

func (s *CostumerRepository) UpdateID(ctx context.Context, id string, input ent.UpdateCostumerInput) (*ent.Costumer, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Costumer.UpdateOneID(uuidId).SetInput(input).Save(ctx)
}
func (s *CostumerRepository) Delete(ctx context.Context) {}
