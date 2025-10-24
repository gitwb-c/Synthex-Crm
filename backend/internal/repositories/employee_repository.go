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

func (s *EmployeeRepository) Read(ctx context.Context) ([]*ent.Employee, error) {
	return s.client.Employee.Query().All(ctx)
}
func (s *EmployeeRepository) ReadID(ctx context.Context, id uuid.UUID) (*ent.Employee, error) {
	return s.client.Employee.Query().Where(employee.IDEQ(id)).Only(ctx)
}

func (s *EmployeeRepository) Create(ctx context.Context, input ent.CreateEmployeeInput) (*ent.Employee, error) {
	return s.client.Employee.Create().SetInput(input).Save(ctx)
}
func (s *EmployeeRepository) UpdateID(ctx context.Context, id string, input ent.UpdateEmployeeInput) (*ent.Employee, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Employee.UpdateOneID(uuidId).SetInput(input).Save(ctx)
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
