package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type DealService struct {
	client *ent.Client
}

func NewDealService() *DealService {
	return &DealService{
		client: db.Client,
	}
}

// ([]*ent.Deal, error)
func (s *DealService) Read(data dto.ReadDealReq) {}

func (s *DealService) Create(data dto.ReadDealReq) {}

func (s *DealService) Update(data dto.ReadDealReq) {}

func (s *DealService) Delete(data dto.ReadDealReq) {}
