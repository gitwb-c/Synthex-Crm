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

func (s *EmployeeAuthRepository) Read(ctx context.Context) ([]*ent.EmployeeAuth, error) {
	return s.client.EmployeeAuth.Query().All(ctx)
}

func (s *EmployeeAuthRepository) ReadEmail(ctx context.Context, email string) (*ent.EmployeeAuth, error) {
	return s.client.EmployeeAuth.Query().Where(employeeauth.Email(email)).WithEmployee(func(q *ent.EmployeeQuery) {
		q.WithTenant().WithDepartment()
	}).Only(ctx)
}

func (s *EmployeeAuthRepository) Create(ctx context.Context, input ent.CreateEmployeeAuthInput) (*ent.EmployeeAuth, error) {
	return s.client.EmployeeAuth.Create().SetInput(input).Save(ctx)
}
func (s *EmployeeAuthRepository) UpdateID(ctx context.Context, id string, input ent.UpdateEmployeeAuthInput) (*ent.EmployeeAuth, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.EmployeeAuth.UpdateOneID(uuidId).SetInput(input).Save(ctx)
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
