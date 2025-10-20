package repositories

import (
	"crm.saas/backend/internal/db"
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
func (s *FormRepository) Read() {}

func (s *FormRepository) Create() {}

func (s *FormRepository) Update() {}

func (s *FormRepository) Delete() {}
