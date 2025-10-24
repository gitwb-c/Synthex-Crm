package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type EmployeeService struct {
	repository *repositories.EmployeeRepository
}

func NewEmployeeService(repository *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repository: repository,
	}
}

func (s *EmployeeService) Read(ctx context.Context) ([]*ent.Employee, error) {
	employee, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeService) Create(ctx context.Context, input ent.CreateEmployeeInput) (*ent.Employee, error) {
	employee, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
func (s *EmployeeService) UpdateID(ctx context.Context, id string, input ent.UpdateEmployeeInput) (*ent.Employee, error) {
	employee, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
func (s *EmployeeService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
