package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
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
func (s *EmployeeRepository) Read(data dto.ReadEmployeeReq) {}

func (s *EmployeeRepository) Create(data dto.CreateEmployeeReq) {}

func (s *EmployeeRepository) Update(data dto.UpdateEmployeeReq) {}

func (s *EmployeeRepository) Delete(data dto.DeleteEmployeeReq) {}
