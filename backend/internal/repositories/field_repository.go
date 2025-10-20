package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
)

type FieldRepository struct {
	client *ent.Client
}

func NewFieldRepository() *FieldRepository {
	return &FieldRepository{
		client: db.Client,
	}
}

// ([]*ent.Field, error)
func (s *FieldRepository) Read() {}

func (s *FieldRepository) Create() {}

func (s *FieldRepository) Update() {}

func (s *FieldRepository) Delete() {}
