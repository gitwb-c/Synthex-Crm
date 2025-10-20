package repositories

import (
	"crm.saas/backend/internal/db"
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
func (s *ChatRepository) Read() {}

func (s *ChatRepository) Create() {}

func (s *ChatRepository) Update() {}

func (s *ChatRepository) Delete() {}
