package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/costumer"
	"github.com/google/uuid"
)

type CostumerRepository struct {
	client *ent.Client
}

func NewCostumerRepository(client *ent.Client) *CostumerRepository {
	return &CostumerRepository{
		client: client,
	}
}

func (s *CostumerRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Costumer, error) {
	return s.client.Costumer.Query().Where(costumer.TenantIdEQ(tenantId)).All(ctx)
}

func (s *CostumerRepository) Create(ctx context.Context, input ent.CreateCostumerInput) (*ent.Costumer, error) {
	return s.client.Costumer.Create().SetInput(input).Save(ctx)
}

func (s *CostumerRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateCostumerInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Costumer.Update().Where(costumer.IDIn(ids...)).SetInput(input).Save(ctx)
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
func (s *CostumerRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Costumer.Delete().Where(costumer.IDIn(ids...)).Exec(ctx)
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
