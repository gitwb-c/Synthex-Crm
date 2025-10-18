package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
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
func (s *FieldRepository) Read(data dto.ReadFieldReq) {}

func (s *FieldRepository) Create(data dto.CreateFieldReq) {}

func (s *FieldRepository) Update(data dto.UpdateFieldReq) {}

func (s *FieldRepository) Delete(data dto.DeleteFieldReq) {}
