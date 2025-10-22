package repositories

import (
	"context"
	"log"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

type CrmFieldRepository struct {
	client *ent.Client
}

func NewCrmFieldRepository(client *ent.Client) *CrmFieldRepository {
	return &CrmFieldRepository{
		client: client,
	}
}

func (s *CrmFieldRepository) Read(ctx context.Context) ([]*ent.CrmField, error) {
	return s.client.CrmField.Query().All(ctx)
}

func (s *CrmFieldRepository) Create(ctx context.Context, input ent.CreateCrmFieldInput) (*ent.CrmField, error) {
	log.Printf("repository: %v", input)

	return s.client.CrmField.Create().SetInput(input).Save(ctx)
}

func (s *CrmFieldRepository) UpdateID(ctx context.Context, id string, input ent.UpdateCrmFieldInput) (*ent.CrmField, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.CrmField.UpdateOneID(uuidId).SetInput(input).Save(ctx)
}

func (s *CrmFieldRepository) Delete(ctx context.Context) {}
