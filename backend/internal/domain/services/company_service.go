package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type CompanyService struct {
	repository *repositories.CompanyRepository
}

func NewCompanyService(repository *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{
		repository: repository,
	}
}

func (s *CompanyService) Read(ctx context.Context) ([]*ent.Company, error) {
	chat, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (s *CompanyService) Create(ctx context.Context, input ent.CreateCompanyInput) (*ent.Company, error) {
	chat, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
func (s *CompanyService) UpdateID(ctx context.Context, id string, input ent.UpdateCompanyInput) (*ent.Company, error) {
	chat, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
func (s *CompanyService) DeleteID(ctx context.Context, id string) error {
	err := s.repository.DeleteID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
