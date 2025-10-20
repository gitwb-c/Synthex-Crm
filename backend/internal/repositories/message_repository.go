package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
)

type MessageRepository struct {
	client *ent.Client
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		client: db.Client,
	}
}

// ([]*ent.Message, error)
func (s *MessageRepository) Read() {}

func (s *MessageRepository) Create() {}

func (s *MessageRepository) Update() {}

func (s *MessageRepository) Delete() {}
