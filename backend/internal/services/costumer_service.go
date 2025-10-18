package services

import (
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/repositories"
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
func (s *CostumerService) Read(data dto.ReadCostumerReq) {}

func (s *CostumerService) Create(data dto.CreateCostumerReq) {}

func (s *CostumerService) Update(data dto.UpdateCostumerReq) {}

func (s *CostumerService) Delete(data dto.DeleteCostumerReq) {}
