package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
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
func (s *DealRepository) Read(data dto.ReadDealReq) {}

func (s *DealRepository) Create(data dto.CreateDealReq) {}

func (s *DealRepository) Update(data dto.UpdateDealReq) {}

func (s *DealRepository) Delete(data dto.DeleteDealReq) {}
