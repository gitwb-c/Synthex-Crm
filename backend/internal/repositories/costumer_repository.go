package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type CostumerRepository struct {
	client *ent.Client
}

func NewCostumerRepository(client *ent.Client) *CostumerRepository {
	return &CostumerRepository{
		client: client,
	}
}

// ([]*ent.Costumer, error)
func (s *CostumerRepository) Read(ctx context.Context) {}

func (s *CostumerRepository) Create(ctx context.Context) {}

func (s *CostumerRepository) Update(ctx context.Context) {}

func (s *CostumerRepository) Delete(ctx context.Context) {}
