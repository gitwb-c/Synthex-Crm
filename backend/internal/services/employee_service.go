package services

import (
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/repositories"
)

type EmployeeService struct {
	repository *repositories.EmployeeRepository
}

func NewEmployeeService(repository *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repository: repository,
	}
}

// ([]*ent.Employee, error)
func (s *EmployeeService) Read(data dto.ReadEmployeeReq) {}

func (s *EmployeeService) Create(data dto.CreateEmployeeReq) {}

func (s *EmployeeService) Update(data dto.UpdateEmployeeReq) {}

func (s *EmployeeService) Delete(data dto.DeleteEmployeeReq) {}
