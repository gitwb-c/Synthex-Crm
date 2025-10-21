package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type EmployeeRepository struct {
	client *ent.Client
}

func NewEmployeeRepository(client *ent.Client) *EmployeeRepository {
	return &EmployeeRepository{
		client: client,
	}
}

// ([]*ent.Employee, error)
func (s *EmployeeRepository) Read(ctx context.Context) {}

func (s *EmployeeRepository) Create(ctx context.Context) {}

func (s *EmployeeRepository) Update(ctx context.Context) {}

func (s *EmployeeRepository) Delete(ctx context.Context) {}
