package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
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
func (s *CompanyRepository) DeleteID(ctx context.Context, id string) error {
	uuidId, e := uuid.Parse(id)
	if e != nil {
		return e
	}
	err := s.client.Company.DeleteOneID(uuidId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
