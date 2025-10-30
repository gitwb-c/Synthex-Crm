package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/google/uuid"
)

type MessageService struct {
	repository *repositories.MessageRepository
}

func NewMessageService(repository *repositories.MessageRepository) *MessageService {
	return &MessageService{
		repository: repository,
	}
}

func (s *MessageService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Message, error) {
	message, err := s.repository.Read(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *MessageService) Create(ctx context.Context, input ent.CreateMessageInput) (*ent.Message, error) {
	message, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return message, nil
}
func (s *MessageService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateMessageInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (s *MessageService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
