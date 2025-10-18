package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type FormRepository struct {
	client *ent.Client
}

func NewFormRepository() *FormRepository {
	return &FormRepository{
		client: db.Client,
	}
}

// ([]*ent.Form, error)
func (s *FormRepository) Read(data dto.ReadFormReq) {}

func (s *FormRepository) Create(data dto.CreateFormReq) {}

func (s *FormRepository) Update(data dto.UpdateFormReq) {}

func (s *FormRepository) Delete(data dto.DeleteFormReq) {}
