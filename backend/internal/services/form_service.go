package services

import (
	"crm.saas/backend/internal/repositories"
)

type FormService struct {
	repository *repositories.FormRepository
}

func NewFormService(repository *repositories.FormRepository) *FormService {
	return &FormService{
		repository: repository,
	}
}

// ([]*ent.Form, error)
func (s *FormService) Read() {}

func (s *FormService) Create() {}

func (s *FormService) Update() {}

func (s *FormService) Delete() {}
