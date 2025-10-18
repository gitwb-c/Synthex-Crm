package services

import (
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/repositories"
)

type FieldService struct {
	repository *repositories.FieldRepository
}

func NewFieldService(repository *repositories.FieldRepository) *FieldService {
	return &FieldService{
		repository: repository,
	}
}

// ([]*ent.Field, error)
func (s *FieldService) Read(data dto.ReadFieldReq) {}

func (s *FieldService) Create(data dto.CreateFieldReq) {}

func (s *FieldService) Update(data dto.UpdateFieldReq) {}

func (s *FieldService) Delete(data dto.DeleteFieldReq) {}
