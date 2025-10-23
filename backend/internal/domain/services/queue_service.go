package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type QueueService struct {
	repository *repositories.QueueRepository
}

func NewQueueService(repository *repositories.QueueRepository) *QueueService {
	return &QueueService{
		repository: repository,
	}
}

func (s *QueueService) Read(ctx context.Context) ([]*ent.Queue, error) {
	queue, err := s.repository.Read(ctx)
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
func (s *QueueService) UpdateID(ctx context.Context, id string, input ent.UpdateQueueInput) (*ent.Queue, error) {
	queue, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return queue, nil
}
func (s *QueueService) DeleteID(ctx context.Context, id string) error {
	err := s.repository.DeleteID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
