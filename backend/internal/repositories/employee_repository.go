package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
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
func (s *EmployeeRepository) Delete(ctx context.Context) {}
