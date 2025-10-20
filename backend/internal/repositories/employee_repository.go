package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
)

type EmployeeRepository struct {
	client *ent.Client
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{
		client: db.Client,
	}
}

// ([]*ent.Employee, error)
func (s *EmployeeRepository) Read() {}

func (s *EmployeeRepository) Create() {}

func (s *EmployeeRepository) Update() {}

func (s *EmployeeRepository) Delete() {}
