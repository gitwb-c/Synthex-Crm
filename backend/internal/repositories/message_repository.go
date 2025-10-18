package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
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
func (s *MessageRepository) Read(data dto.ReadMessageReq) {}

func (s *MessageRepository) Create(data dto.CreateMessageReq) {}

func (s *MessageRepository) Update(data dto.UpdateMessageReq) {}

func (s *MessageRepository) Delete(data dto.DeleteMessageReq) {}
