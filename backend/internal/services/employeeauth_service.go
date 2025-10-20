package services

import (
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
func (s *EmployeeAuthService) Read() {}

func (s *EmployeeAuthService) Create() {}

func (s *EmployeeAuthService) Update() {}

func (s *EmployeeAuthService) Delete() {}
