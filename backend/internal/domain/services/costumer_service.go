package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type CostumerService struct {
	repository *repositories.CostumerRepository
}

func NewCostumerService(repository *repositories.CostumerRepository) *CostumerService {
	return &CostumerService{
		repository: repository,
	}
}

func (s *CostumerService) Read(ctx context.Context) ([]*ent.Costumer, error) {
	costumer, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return costumer, nil
}
func (s *CostumerService) Create(ctx context.Context, input ent.CreateCostumerInput) (*ent.Costumer, error) {
	costumer, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return costumer, nil
}

func (s *CostumerService) UpdateID(ctx context.Context, id string, input ent.UpdateCostumerInput) (*ent.Costumer, error) {
	costumer, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return costumer, nil
}

func (s *CostumerService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
