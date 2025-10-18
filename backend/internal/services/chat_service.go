package services

import (
	"crm.saas/backend/internal/dto"
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
func (s *ChatService) Read(data dto.ReadChatReq) {}

func (s *ChatService) Create(data dto.ReadChatReq) {}

func (s *ChatService) Update(data dto.ReadChatReq) {}

func (s *ChatService) Delete(data dto.ReadChatReq) {}
