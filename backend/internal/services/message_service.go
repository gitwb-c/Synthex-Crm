package services

import (
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/repositories"
)

type MessageService struct {
	repository *repositories.MessageRepository
}

func NewMessageService(repository *repositories.MessageRepository) *MessageService {
	return &MessageService{
		repository: repository,
	}
}

// ([]*ent.Message, error)
func (s *MessageService) Read(data dto.ReadMessageReq) {}

func (s *MessageService) Create(data dto.CreateMessageReq) {}

func (s *MessageService) Update(data dto.UpdateMessageReq) {}

func (s *MessageService) Delete(data dto.DeleteMessageReq) {}
