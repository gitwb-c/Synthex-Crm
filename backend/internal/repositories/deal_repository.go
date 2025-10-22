package repositories

import (
	"context"
	"log"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

type DealRepository struct {
	client *ent.Client
}

func NewDealRepository(client *ent.Client) *DealRepository {
	return &DealRepository{
		client: client,
	}
}

func (s *DealRepository) Read(ctx context.Context) ([]*ent.Deal, error) {
	return s.client.Deal.Query().All(ctx)
}

func (s *DealRepository) Create(ctx context.Context, input ent.CreateDealInput) (*ent.Deal, error) {
	log.Printf("repository: %v", input)

	return s.client.Deal.Create().SetInput(input).Save(ctx)
}

func (s *DealRepository) UpdateID(ctx context.Context, id string, input ent.UpdateDealInput) (*ent.Deal, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Deal.UpdateOneID(uuidId).SetInput(input).Save(ctx)
}

func (s *DealRepository) Delete(ctx context.Context) {}
