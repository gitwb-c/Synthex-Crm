package services

import (
	"crm.saas/backend/internal/dto"
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
func (s *FormService) Read(data dto.ReadFormReq) {}

func (s *FormService) Create(data dto.CreateFormReq) {}

func (s *FormService) Update(data dto.UpdateFormReq) {}

func (s *FormService) Delete(data dto.DeleteFormReq) {}
