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

func (s *DealRepository) Read(ctx context.Context) ([]*ent.Deal, error) {
	return s.client.Deal.Query().All(ctx)
}

func (s *DealRepository) Create(ctx context.Context, input ent.CreateDealInput) (*ent.Deal, error) {
	return s.client.Deal.Create().SetInput(input).Save(ctx)
}

func (s *DealRepository) UpdateID(ctx context.Context, id string, input ent.UpdateDealInput) (*ent.Deal, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Deal.UpdateOneID(uuidId).SetInput(input).Save(ctx)
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
