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

func (s *RbacRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Rbac, error) {
	return s.client.Rbac.Query().Where(rbac.TenantIdEQ(tenantId)).All(ctx)
}

func (s *RbacRepository) Create(ctx context.Context, input ent.CreateRbacInput) (*ent.Rbac, error) {
	return s.client.Rbac.Create().SetInput(input).Save(ctx)
}
func (s *RbacRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateRbacInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Rbac.Update().Where(rbac.IDIn(ids...)).SetInput(input).Save(ctx)
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
