package services

import (
	"context"
	"log"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type DealService struct {
	repository *repositories.DealRepository
}

func NewDealService(repository *repositories.DealRepository) *DealService {
	return &DealService{
		repository: repository,
	}
}

func (s *DealService) Read(ctx context.Context) ([]*ent.Deal, error) {
	deal, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return deal, nil
}

func (s *DealService) Create(ctx context.Context, input ent.CreateDealInput) (*ent.Deal, error) {
	log.Printf("service: %v", input)
	deal, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return deal, nil
}

func (s *DealService) UpdateID(ctx context.Context, id string, input ent.UpdateDealInput) (*ent.Deal, error) {
	deal, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return deal, nil
}

func (s *DealService) Delete() {}
