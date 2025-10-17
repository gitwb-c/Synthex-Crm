package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type FieldService struct {
	client *ent.Client
}

func NewFieldService() *FieldService {
	return &FieldService{
		client: db.Client,
	}
}

// ([]*ent.Field, error)
func (s *FieldService) Read(data dto.ReadFieldReq) {}

func (s *FieldService) Create(data dto.ReadFieldReq) {}

func (s *FieldService) Update(data dto.ReadFieldReq) {}

func (s *FieldService) Delete(data dto.ReadFieldReq) {}
