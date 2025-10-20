package services

import (
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
func (s *DealService) Read() {}

func (s *DealService) Create() {}

func (s *DealService) Update() {}

func (s *DealService) Delete() {}
