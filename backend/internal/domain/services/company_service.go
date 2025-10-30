package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type CompanyService struct {
	repository *repositories.CompanyRepository
}

func NewCompanyService(repository *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{
		repository: repository,
	}
}

func (s *CompanyService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Company, error) {
	company, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (s *CompanyService) Create(ctx context.Context, input ent.CreateCompanyInput, client *ent.Client) (*ent.Company, error) {
	company, err := s.repository.Create(ctx, input, client)
	if err != nil {
		return nil, err
	}
	return company, nil
}
func (s *CompanyService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateCompanyInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}
func (s *CompanyService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
