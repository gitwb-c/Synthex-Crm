package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type DepartmentRepository struct {
	client *ent.Client
}

func NewDepartmentRepository(client *ent.Client) *DepartmentRepository {
	return &DepartmentRepository{
		client: client,
	}
}

// ([]*ent.Department, error)
func (s *DepartmentRepository) Read(ctx context.Context) {}

func (s *DepartmentRepository) Create(ctx context.Context) {}

func (s *DepartmentRepository) Update(ctx context.Context) {}

func (s *DepartmentRepository) Delete(ctx context.Context) {}
