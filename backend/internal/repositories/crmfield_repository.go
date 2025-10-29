package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/crmfield"
	"github.com/google/uuid"
)

type CrmFieldRepository struct {
	client *ent.Client
}

func NewCrmFieldRepository(client *ent.Client) *CrmFieldRepository {
	return &CrmFieldRepository{
		client: client,
	}
}

func (s *CrmFieldRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.CrmField, error) {
	return s.client.CrmField.Query().Where(crmfield.TenantIdEQ(tenantId)).All(ctx)
}

func (s *CrmFieldRepository) Create(ctx context.Context, input ent.CreateCrmFieldInput) (*ent.CrmField, error) {
	return s.client.CrmField.Create().SetInput(input).Save(ctx)
}

func (s *CrmFieldRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateCrmFieldInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.CrmField.Update().Where(crmfield.IDIn(ids...)).SetInput(input).Save(ctx)
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

func (s *CrmFieldRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.CrmField.Delete().Where(crmfield.IDIn(ids...)).Exec(ctx)
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
