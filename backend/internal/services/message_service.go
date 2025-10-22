package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type MessageService struct {
	repository *repositories.MessageRepository
}

func NewMessageService(repository *repositories.MessageRepository) *MessageService {
	return &MessageService{
		repository: repository,
	}
}

func (s *MessageService) Read(ctx context.Context) ([]*ent.Message, error) {
	message, err := s.repository.Read(ctx)
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
func (s *MessageService) UpdateID(ctx context.Context, id string, input ent.UpdateMessageInput) (*ent.Message, error) {
	message, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *MessageService) Delete() {}
