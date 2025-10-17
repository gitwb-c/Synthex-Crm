package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type ChatService struct {
	client *ent.Client
}

func NewChatService() *ChatService {
	return &ChatService{
		client: db.Client,
	}
}

// ([]*ent.Chat, error)
func (s *ChatService) Read(data dto.ReadChatReq) {}

func (s *ChatService) Create(data dto.ReadChatReq) {}

func (s *ChatService) Update(data dto.ReadChatReq) {}

func (s *ChatService) Delete(data dto.ReadChatReq) {}
