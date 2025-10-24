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

func (s *RbacService) Read(ctx context.Context) ([]*ent.Rbac, error) {
	chat, err := s.repository.Read(ctx)
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
func (s *RbacService) UpdateID(ctx context.Context, id string, input ent.UpdateRbacInput) (*ent.Rbac, error) {
	chat, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
func (s *RbacService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
