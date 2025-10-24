package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/rbac"
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
func (s *RbacRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Rbac.Delete().Where(rbac.IDIn(ids...)).Exec(ctx)
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
