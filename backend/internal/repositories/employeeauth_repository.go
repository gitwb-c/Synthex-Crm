package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/employeeauth"
	"github.com/google/uuid"
)

type EmployeeAuthRepository struct {
	client *ent.Client
}

func NewEmployeeAuthRepository(client *ent.Client) *EmployeeAuthRepository {
	return &EmployeeAuthRepository{
		client: client,
	}
}

func (s *EmployeeAuthRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.EmployeeAuth, error) {
	return s.client.EmployeeAuth.Query().Where(employeeauth.TenantIdEQ(tenantId)).All(ctx)
}

func (s *EmployeeAuthRepository) ReadEmail(ctx context.Context, email string, tenantId uuid.UUID) (*ent.EmployeeAuth, error) {
	return s.client.EmployeeAuth.Query().Where(employeeauth.Email(email), employeeauth.TenantIdEQ(tenantId)).WithEmployee(func(q *ent.EmployeeQuery) {
		q.WithTenant().WithDepartment()
	}).WithTenant().Only(ctx)
}

func (s *EmployeeAuthRepository) Create(ctx context.Context, input ent.CreateEmployeeAuthInput, client *ent.Client) (*ent.EmployeeAuth, error) {
	if client != nil {
		return client.EmployeeAuth.Create().SetInput(input).Save(ctx)
	}
	return s.client.EmployeeAuth.Create().SetInput(input).Save(ctx)
}
func (s *EmployeeAuthRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateEmployeeAuthInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.EmployeeAuth.Update().Where(employeeauth.IDIn(ids...)).SetInput(input).Save(ctx)
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
func (s *EmployeeAuthRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.EmployeeAuth.Delete().Where(employeeauth.IDIn(ids...)).Exec(ctx)
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
