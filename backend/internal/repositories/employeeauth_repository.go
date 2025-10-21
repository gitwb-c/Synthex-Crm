package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type EmployeeAuthRepository struct {
	client *ent.Client
}

func NewEmployeeAuthRepository(client *ent.Client) *EmployeeAuthRepository {
	return &EmployeeAuthRepository{
		client: client,
	}
}

// ([]*ent.EmployeeAuth, error)
func (s *EmployeeAuthRepository) Read(ctx context.Context) {}

func (s *EmployeeAuthRepository) Create(ctx context.Context) {}

func (s *EmployeeAuthRepository) Update(ctx context.Context) {}

func (s *EmployeeAuthRepository) Delete(ctx context.Context) {}
