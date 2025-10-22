package repositories

import (
	"context"
	"log"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

type DealCrmFieldRepository struct {
	client *ent.Client
}

func NewDealCrmFieldRepository(client *ent.Client) *DealCrmFieldRepository {
	return &DealCrmFieldRepository{
		client: client,
	}
}

func (s *DealCrmFieldRepository) Read(ctx context.Context) ([]*ent.DealCrmField, error) {
	return s.client.DealCrmField.Query().All(ctx)
}

func (s *DealCrmFieldRepository) Create(ctx context.Context, input ent.CreateDealCrmFieldInput) (*ent.DealCrmField, error) {
	log.Printf("repository: %v", input)

	return s.client.DealCrmField.Create().SetInput(input).Save(ctx)
}

func (s *DealCrmFieldRepository) UpdateID(ctx context.Context, id string, input ent.UpdateDealCrmFieldInput) (*ent.DealCrmField, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.DealCrmField.UpdateOneID(uuidId).SetInput(input).Save(ctx)
}

func (s *DealCrmFieldRepository) Delete(ctx context.Context) {}
