package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type CostumerService struct {
	client *ent.Client
}

func NewCostumerService() *CostumerService {
	return &CostumerService{
		client: db.Client,
	}
}

// ([]*ent.Costumer, error)
func (s *CostumerService) Read(data dto.ReadCostumerReq) {}

func (s *CostumerService) Create(data dto.ReadCostumerReq) {}

func (s *CostumerService) Update(data dto.ReadCostumerReq) {}

func (s *CostumerService) Delete(data dto.ReadCostumerReq) {}
