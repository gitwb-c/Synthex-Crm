package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/company"
	"github.com/google/uuid"
)

type CompanyRepository struct {
	client *ent.Client
}

func NewCompanyRepository(client *ent.Client) *CompanyRepository {
	return &CompanyRepository{
		client: client,
	}
}

func (s *CompanyRepository) Read(ctx context.Context) ([]*ent.Company, error) {
	return s.client.Company.Query().All(ctx)
}

func (s *CompanyRepository) Create(ctx context.Context, input ent.CreateCompanyInput, client *ent.Client) (*ent.Company, error) {
	if client != nil {
		return client.Company.Create().SetInput(input).Save(ctx)
	}
	return s.client.Company.Create().SetInput(input).Save(ctx)
}
func (s *CompanyRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateCompanyInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Company.Update().Where(company.IDIn(ids...)).SetInput(input).Save(ctx)
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
func (s *CompanyRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Company.Delete().Where(company.IDIn(ids...)).Exec(ctx)
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
