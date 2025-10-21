package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/graphql"
)

type DealRepository struct {
	client *ent.Client
}

func NewDealRepository(client *ent.Client) *DealRepository {
	return &DealRepository{
		client: client,
	}
}

// ([]*ent.Deal, error)
func (s *DealRepository) Read(ctx context.Context) {}

func (s *DealRepository) Create(ctx context.Context, input *graphql.CreateDealInput) (*ent.Deal, error) {
	return s.client.Deal.Create().SetInput(input).Save(ctx)
}

func (s *DealRepository) Update(ctx context.Context) {}

func (s *DealRepository) Delete(ctx context.Context) {}
