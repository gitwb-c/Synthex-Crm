package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
)

type DepartmentRepository struct {
	client *ent.Client
}

func NewDepartmentRepository() *DepartmentRepository {
	return &DepartmentRepository{
		client: db.Client,
	}
}

// ([]*ent.Department, error)
func (s *DepartmentRepository) Read() {}

func (s *DepartmentRepository) Create() {}

func (s *DepartmentRepository) Update() {}

func (s *DepartmentRepository) Delete() {}
