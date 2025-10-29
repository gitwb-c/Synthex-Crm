package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/dealcrmfield"
	"github.com/google/uuid"
)

type DealCrmFieldRepository struct {
	client *ent.Client
}

func NewDealCrmFieldRepository(client *ent.Client) *DealCrmFieldRepository {
	return &DealCrmFieldRepository{
		client: client,
	}
}

func (s *DealCrmFieldRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.DealCrmField, error) {
	return s.client.DealCrmField.Query().Where(dealcrmfield.TenantIdEQ(tenantId)).All(ctx)
}

func (s *DealCrmFieldRepository) Create(ctx context.Context, input ent.CreateDealCrmFieldInput) (*ent.DealCrmField, error) {

	return s.client.DealCrmField.Create().SetInput(input).Save(ctx)
}

func (s *DealCrmFieldRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateDealCrmFieldInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.DealCrmField.Update().Where(dealcrmfield.IDIn(ids...)).SetInput(input).Save(ctx)
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

func (s *DealCrmFieldRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.DealCrmField.Delete().Where(dealcrmfield.IDIn(ids...)).Exec(ctx)
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
