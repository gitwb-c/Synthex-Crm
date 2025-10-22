package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type DealCrmFieldService struct {
	repository *repositories.DealCrmFieldRepository
}

func NewDealCrmFieldService(repository *repositories.DealCrmFieldRepository) *DealCrmFieldService {
	return &DealCrmFieldService{
		repository: repository,
	}
}

func (s *DealCrmFieldService) Read(ctx context.Context) ([]*ent.DealCrmField, error) {
	dealCrmField, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return dealCrmField, nil
}

func (s *DealCrmFieldService) Create(ctx context.Context, input ent.CreateDealCrmFieldInput) (*ent.DealCrmField, error) {
	dealCrmField, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return dealCrmField, nil
}
func (s *DealCrmFieldService) UpdateID(ctx context.Context, id string, input ent.UpdateDealCrmFieldInput) (*ent.DealCrmField, error) {
	dealCrmField, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return dealCrmField, nil
}
func (s *DealCrmFieldService) DeleteID(ctx context.Context, id string) error {
	err := s.repository.DeleteID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
