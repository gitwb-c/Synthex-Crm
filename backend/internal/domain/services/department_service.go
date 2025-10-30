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

func (s *DepartmentService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Department, error) {
	department, err := s.repository.Read(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	return department, nil
}
func (s *DepartmentService) ReadRbacs(ctx context.Context, uuid uuid.UUID) (*ent.Department, error) {
	department, err := s.repository.ReadRbacs(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (s *DepartmentService) Create(ctx context.Context, input ent.CreateDepartmentInput, client *ent.Client) (*ent.Department, error) {
	department, err := s.repository.Create(ctx, input, client)
	if err != nil {
		return nil, err
	}
	return department, nil
}
func (s *DepartmentService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateDepartmentInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (s *DepartmentService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
