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

func (s *CompanyRepository) Create(ctx context.Context, input ent.CreateCompanyInput) (*ent.Company, error) {
	return s.client.Company.Create().SetInput(input).Save(ctx)
}
func (s *CompanyRepository) UpdateID(ctx context.Context, id string, input ent.UpdateCompanyInput) (*ent.Company, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Company.UpdateOneID(uuidId).SetInput(input).Save(ctx)
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
