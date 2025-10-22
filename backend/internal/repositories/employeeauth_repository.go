package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
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
func (s *EmployeeAuthRepository) DeleteID(ctx context.Context, id string) error {
	uuidId, e := uuid.Parse(id)
	if e != nil {
		return e
	}
	err := s.client.EmployeeAuth.DeleteOneID(uuidId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
