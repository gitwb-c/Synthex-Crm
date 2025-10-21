package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/graphql"
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

// ([]*ent.Deal, error)
func (s *DealService) Read(ctx context.Context) {}

func (s *DealService) Create(ctx context.Context, input *graphql.CreateDealInput) (*ent.Deal, error) {
	createdDeal, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return createdDeal, nil

}

func (s *DealService) Update(ctx context.Context) {}

func (s *DealService) Delete() {}
