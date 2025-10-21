package services

import (
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type CostumerService struct {
	repository *repositories.CostumerRepository
}

func NewCostumerService(repository *repositories.CostumerRepository) *CostumerService {
	return &CostumerService{
		repository: repository,
	}
}

// ([]*ent.Costumer, error)
func (s *CostumerService) Read() {}

func (s *CostumerService) Create() {}

func (s *CostumerService) Update() {}

func (s *CostumerService) Delete() {}
