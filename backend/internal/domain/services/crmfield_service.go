package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type CrmFieldService struct {
	repository *repositories.CrmFieldRepository
}

func NewCrmFieldService(repository *repositories.CrmFieldRepository) *CrmFieldService {
	return &CrmFieldService{
		repository: repository,
	}
}

func (s *CrmFieldService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.CrmField, error) {
	crmField, err := s.repository.Read(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	return crmField, nil
}

func (s *CrmFieldService) Create(ctx context.Context, input ent.CreateCrmFieldInput) (*ent.CrmField, error) {
	crmField, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return crmField, nil
}
func (s *CrmFieldService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateCrmFieldInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}
func (s *CrmFieldService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
