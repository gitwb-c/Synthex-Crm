package services

import (
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/repositories"
)

type DealService struct {
	repository *repositories.DealRepository
}

func NewDealService(repository *repositories.DealRepository) *DealService {
	return &DealService{
		repository: repository,
	}
}

// ([]*ent.Deal, error)
func (s *DealService) Read(data dto.ReadDealReq) {}

func (s *DealService) Create(data dto.CreateDealReq) {}

func (s *DealService) Update(data dto.UpdateDealReq) {}

func (s *DealService) Delete(data dto.DeleteDealReq) {}
