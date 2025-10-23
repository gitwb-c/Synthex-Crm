package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

type RbacRepository struct {
	client *ent.Client
}

func NewRbacRepository(client *ent.Client) *RbacRepository {
	return &RbacRepository{
		client: client,
	}
}

func (s *RbacRepository) Read(ctx context.Context) ([]*ent.Rbac, error) {
	return s.client.Rbac.Query().All(ctx)
}

func (s *RbacRepository) Create(ctx context.Context, input ent.CreateRbacInput) (*ent.Rbac, error) {
	return s.client.Rbac.Create().SetInput(input).Save(ctx)
}
func (s *RbacRepository) UpdateID(ctx context.Context, id string, input ent.UpdateRbacInput) (*ent.Rbac, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Rbac.UpdateOneID(uuidId).SetInput(input).Save(ctx)
}
func (s *RbacRepository) DeleteID(ctx context.Context, id string) error {
	uuidId, e := uuid.Parse(id)
	if e != nil {
		return e
	}
	err := s.client.Rbac.DeleteOneID(uuidId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
