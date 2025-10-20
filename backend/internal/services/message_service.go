package services

import (
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
func (s *MessageService) Read() {}

func (s *MessageService) Create() {}

func (s *MessageService) Update() {}

func (s *MessageService) Delete() {}
