package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/deal"
	"github.com/google/uuid"
)

type DealRepository struct {
	client *ent.Client
}

func NewDealRepository(client *ent.Client) *DealRepository {
	return &DealRepository{
		client: client,
	}
}

func (s *DealRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Deal, error) {
	return s.client.Deal.Query().Where(deal.TenantIdEQ(tenantId)).All(ctx)
}

func (s *DealRepository) Create(ctx context.Context, input ent.CreateDealInput) (*ent.Deal, error) {
	return s.client.Deal.Create().SetInput(input).Save(ctx)
}

func (s *DealRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateDealInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Deal.Update().Where(deal.IDIn(ids...)).SetInput(input).Save(ctx)
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

func (s *DealRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Deal.Delete().Where(deal.IDIn(ids...)).Exec(ctx)
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
