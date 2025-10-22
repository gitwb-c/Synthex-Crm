package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type EmployeeAuthService struct {
	repository *repositories.EmployeeAuthRepository
}

func NewEmployeeAuthService(repository *repositories.EmployeeAuthRepository) *EmployeeAuthService {
	return &EmployeeAuthService{
		repository: repository,
	}
}

func (s *EmployeeAuthService) Read(ctx context.Context) ([]*ent.EmployeeAuth, error) {
	employeeAuth, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return employeeAuth, nil
}

func (s *EmployeeAuthService) Create(ctx context.Context, input ent.CreateEmployeeAuthInput) (*ent.EmployeeAuth, error) {
	employeeAuth, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return employeeAuth, nil
}
func (s *EmployeeAuthService) UpdateID(ctx context.Context, id string, input ent.UpdateEmployeeAuthInput) (*ent.EmployeeAuth, error) {
	employeeAuth, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return employeeAuth, nil
}

func (s *EmployeeAuthService) DeleteID(ctx context.Context, id string) error {
	err := s.repository.DeleteID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
