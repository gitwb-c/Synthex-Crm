package services

import (
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
func (s *FieldService) Read() {}

func (s *FieldService) Create() {}

func (s *FieldService) Update() {}

func (s *FieldService) Delete() {}
