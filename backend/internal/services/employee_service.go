package services

import (
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
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
func (s *EmployeeService) Read() {}

func (s *EmployeeService) Create() {}

func (s *EmployeeService) Update() {}

func (s *EmployeeService) Delete() {}
