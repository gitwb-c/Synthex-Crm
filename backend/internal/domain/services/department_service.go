package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type DepartmentService struct {
	repository *repositories.DepartmentRepository
}

func NewDepartmentService(repository *repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		repository: repository,
	}
}

func (s *DepartmentService) Read(ctx context.Context) ([]*ent.Department, error) {
	department, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (s *DepartmentService) Create(ctx context.Context, input ent.CreateDepartmentInput) (*ent.Department, error) {
	department, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return department, nil
}
func (s *DepartmentService) UpdateID(ctx context.Context, id string, input ent.UpdateDepartmentInput) (*ent.Department, error) {
	department, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (s *DepartmentService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
