package services

import (
	"crm.saas/backend/internal/repositories"
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
func (s *ChatService) Read() {}

func (s *ChatService) Create() {}

func (s *ChatService) Update() {}

func (s *ChatService) Delete() {}
