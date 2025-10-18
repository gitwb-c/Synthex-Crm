package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
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
func (s *CostumerRepository) Read(data dto.ReadCostumerReq) {}

func (s *CostumerRepository) Create(data dto.CreateCostumerReq) {}

func (s *CostumerRepository) Update(data dto.UpdateCostumerReq) {}

func (s *CostumerRepository) Delete(data dto.DeleteCostumerReq) {}
