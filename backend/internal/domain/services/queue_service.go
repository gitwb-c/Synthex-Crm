package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type QueueService struct {
	repository *repositories.QueueRepository
}

func NewQueueService(repository *repositories.QueueRepository) *QueueService {
	return &QueueService{
		repository: repository,
	}
}

func (s *QueueService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Queue, error) {
	queue, err := s.repository.Read(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	return queue, nil
}

func (s *QueueService) Create(ctx context.Context, input ent.CreateQueueInput) (*ent.Queue, error) {
	queue, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return queue, nil
}
func (s *QueueService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateQueueInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}
func (s *QueueService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
