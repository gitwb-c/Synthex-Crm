package services

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type ChatService struct {
	repository *repositories.ChatRepository
}

func NewChatService(repository *repositories.ChatRepository) *ChatService {
	return &ChatService{
		repository: repository,
	}
}

// ([]*ent.Chat, error)
func (s *ChatService) Read(ctx context.Context) ([]*ent.Chat, error) {
	chat, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (s *ChatService) Create(ctx context.Context, input ent.CreateChatInput) (*ent.Chat, error) {
	chat, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
func (s *ChatService) UpdateID(ctx context.Context, id string, input ent.UpdateChatInput) (*ent.Chat, error) {
	chat, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
func (s *ChatService) Delete() {}
