package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
)

type DealRepository struct {
	client *ent.Client
}

func NewDealRepository() *DealRepository {
	return &DealRepository{
		client: db.Client,
	}
}

// ([]*ent.Deal, error)
func (s *DealRepository) Read() {}

func (s *DealRepository) Create() {}

func (s *DealRepository) Update() {}

func (s *DealRepository) Delete() {}
