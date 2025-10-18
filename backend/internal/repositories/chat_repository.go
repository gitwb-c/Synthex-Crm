package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type ChatRepository struct {
	client *ent.Client
}

func NewChatRepository() *ChatRepository {
	return &ChatRepository{
		client: db.Client,
	}
}

// ([]*ent.Chat, error)
func (s *ChatRepository) Read(data dto.ReadChatReq) {}

func (s *ChatRepository) Create(data dto.CreateChatReq) {}

func (s *ChatRepository) Update(data dto.UpdateChatReq) {}

func (s *ChatRepository) Delete(data dto.DeleteChatReq) {}
