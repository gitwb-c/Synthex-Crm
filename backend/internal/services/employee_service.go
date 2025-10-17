package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type EmployeeService struct {
	client *ent.Client
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		client: db.Client,
	}
}

// ([]*ent.Employee, error)
func (s *EmployeeService) Read(data dto.ReadEmployeeReq) {}

func (s *EmployeeService) Create(data dto.ReadEmployeeReq) {}

func (s *EmployeeService) Update(data dto.ReadEmployeeReq) {}

func (s *EmployeeService) Delete(data dto.ReadEmployeeReq) {}
