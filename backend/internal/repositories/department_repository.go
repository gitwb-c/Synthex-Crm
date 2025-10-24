package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/department"
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
func (s *DepartmentRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Department.Delete().Where(department.IDIn(ids...)).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
