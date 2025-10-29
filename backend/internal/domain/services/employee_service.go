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

func (s *EmployeeService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Employee, error) {
	employee, err := s.repository.Read(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeService) ReadID(ctx context.Context, id uuid.UUID) (*ent.Employee, error) {
	employee, err := s.repository.ReadID(ctx, id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeService) Create(ctx context.Context, input ent.CreateEmployeeInput, client *ent.Client) (*ent.Employee, error) {
	employee, err := s.repository.Create(ctx, input, client)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
func (s *EmployeeService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateEmployeeInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}
func (s *EmployeeService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
