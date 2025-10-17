package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type FormService struct {
	client *ent.Client
}

func NewFormService() *FormService {
	return &FormService{
		client: db.Client,
	}
}

// ([]*ent.Form, error)
func (s *FormService) Read(data dto.ReadFormReq) {}

func (s *FormService) Create(data dto.ReadFormReq) {}

func (s *FormService) Update(data dto.ReadFormReq) {}

func (s *FormService) Delete(data dto.ReadFormReq) {}
