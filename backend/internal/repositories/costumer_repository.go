package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
)

type CostumerRepository struct {
	client *ent.Client
}

func NewCostumerRepository() *CostumerRepository {
	return &CostumerRepository{
		client: db.Client,
	}
}

// ([]*ent.Costumer, error)
func (s *CostumerRepository) Read() {}

func (s *CostumerRepository) Create() {}

func (s *CostumerRepository) Update() {}

func (s *CostumerRepository) Delete() {}
