package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/employee"
	"github.com/google/uuid"
)

type EmployeeRepository struct {
	client *ent.Client
}

func NewEmployeeRepository(client *ent.Client) *EmployeeRepository {
	return &EmployeeRepository{
		client: client,
	}
}

func (s *EmployeeRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Employee, error) {
	return s.client.Employee.Query().Where(employee.TenantIdEQ(tenantId)).All(ctx)
}
func (s *EmployeeRepository) ReadID(ctx context.Context, id uuid.UUID) (*ent.Employee, error) {
	return s.client.Employee.Query().Where(employee.IDEQ(id)).Only(ctx)
}

func (s *EmployeeRepository) Create(ctx context.Context, input ent.CreateEmployeeInput, client *ent.Client) (*ent.Employee, error) {
	if client != nil {
		return client.Employee.Create().SetInput(input).Save(ctx)
	}
	return s.client.Employee.Create().SetInput(input).Save(ctx)
}
func (s *EmployeeRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateEmployeeInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Employee.Update().Where(employee.IDIn(ids...)).SetInput(input).Save(ctx)
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
func (s *EmployeeRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Employee.Delete().Where(employee.IDIn(ids...)).Exec(ctx)
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
