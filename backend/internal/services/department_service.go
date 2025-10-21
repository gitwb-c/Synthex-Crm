package services

import (
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type DepartmentService struct {
	repository *repositories.DepartmentRepository
}

func NewDepartmentService(repository *repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		repository: repository,
	}
}

// ([]*ent.Department, error)
func (s *DepartmentService) Read() {}

func (s *DepartmentService) Create() {}

func (s *DepartmentService) Update() {}

func (s *DepartmentService) Delete() {}
