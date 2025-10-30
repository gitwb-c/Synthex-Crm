package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type DealService struct {
	repository *repositories.DealRepository
}

func NewDealService(repository *repositories.DealRepository) *DealService {
	return &DealService{
		repository: repository,
	}
}

func (s *DealService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Deal, error) {
	deal, err := s.repository.Read(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	return deal, nil
}

func (s *DealService) Create(ctx context.Context, input ent.CreateDealInput) (*ent.Deal, error) {
	deal, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return deal, nil
}

func (s *DealService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateDealInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (s *DealService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
