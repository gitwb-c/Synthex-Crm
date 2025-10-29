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

func (s *DepartmentRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Department, error) {
	return s.client.Department.Query().Where(department.TenantIdEQ(tenantId)).All(ctx)
}
func (s *DepartmentRepository) ReadRbacs(ctx context.Context, id uuid.UUID) (*ent.Department, error) {
	return s.client.Department.Query().Where(department.IDEQ(id)).WithRbacs().Only(ctx)
}

func (s *DepartmentRepository) Create(ctx context.Context, input ent.CreateDepartmentInput, client *ent.Client) (*ent.Department, error) {
	if client != nil {
		return client.Department.Create().SetInput(input).Save(ctx)
	}
	return s.client.Department.Create().SetInput(input).Save(ctx)
}
func (s *DepartmentRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateDepartmentInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Department.Update().Where(department.IDIn(ids...)).SetInput(input).Save(ctx)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	return n, nil
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
