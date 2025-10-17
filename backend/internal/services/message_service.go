package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type MessageService struct {
	client *ent.Client
}

func NewMessageService() *MessageService {
	return &MessageService{
		client: db.Client,
	}
}

// ([]*ent.Message, error)
func (s *MessageService) Read(data dto.ReadMessageReq) {}

func (s *MessageService) Create(data dto.ReadMessageReq) {}

func (s *MessageService) Update(data dto.ReadMessageReq) {}

func (s *MessageService) Delete(data dto.ReadMessageReq) {}
