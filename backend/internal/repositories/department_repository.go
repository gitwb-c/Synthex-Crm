package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

type DepartmentRepository struct {
	client *ent.Client
}

func NewDepartmentRepository(client *ent.Client) *DepartmentRepository {
	return &DepartmentRepository{
		client: client,
	}
}

func (s *DepartmentRepository) Read(ctx context.Context) ([]*ent.Department, error) {
	return s.client.Department.Query().All(ctx)
}

func (s *DepartmentRepository) Create(ctx context.Context, input ent.CreateDepartmentInput) (*ent.Department, error) {
	return s.client.Department.Create().SetInput(input).Save(ctx)
}
func (s *DepartmentRepository) UpdateID(ctx context.Context, id string, input ent.UpdateDepartmentInput) (*ent.Department, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Department.UpdateOneID(uuidId).SetInput(input).Save(ctx)
}
func (s *DepartmentRepository) DeleteID(ctx context.Context, id string) error {
	uuidId, e := uuid.Parse(id)
	if e != nil {
		return e
	}
	err := s.client.Department.DeleteOneID(uuidId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
