package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type RbacService struct {
	repository *repositories.RbacRepository
}

func NewRbacService(repository *repositories.RbacRepository) *RbacService {
	return &RbacService{
		repository: repository,
	}
}

func (s *RbacService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Rbac, error) {
	chat, err := s.repository.Read(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (s *RbacService) Create(ctx context.Context, input ent.CreateRbacInput) (*ent.Rbac, error) {
	chat, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
func (s *RbacService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateRbacInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}
func (s *RbacService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
