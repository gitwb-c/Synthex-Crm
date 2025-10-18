package services

import (
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/repositories"
)

type EmployeeAuthService struct {
	repository *repositories.EmployeeAuthRepository
}

func NewEmployeeAuthService(repository *repositories.EmployeeAuthRepository) *EmployeeAuthService {
	return &EmployeeAuthService{
		repository: repository,
	}
}

// ([]*ent.EmployeeAuth, error)
func (s *EmployeeAuthService) Read(data dto.ReadEmployeeAuthReq) {}

func (s *EmployeeAuthService) Create(data dto.CreateEmployeeAuthReq) {}

func (s *EmployeeAuthService) Update(data dto.UpdateEmployeeAuthReq) {}

func (s *EmployeeAuthService) Delete(data dto.DeleteEmployeeAuthReq) {}
