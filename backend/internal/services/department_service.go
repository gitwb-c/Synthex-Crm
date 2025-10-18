package services

import (
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/repositories"
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
func (s *DepartmentService) Read(data dto.ReadDepartmentReq) {}

func (s *DepartmentService) Create(data dto.CreateDepartmentReq) {}

func (s *DepartmentService) Update(data dto.UpdateDepartmentReq) {}

func (s *DepartmentService) Delete(data dto.DeleteDepartmentReq) {}
