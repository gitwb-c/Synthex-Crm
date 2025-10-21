package services

import (
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
)

type QueueService struct {
	repository *repositories.QueueRepository
}

func NewQueueService(repository *repositories.QueueRepository) *QueueService {
	return &QueueService{
		repository: repository,
	}
}

// ([]*ent.Queue, error)
func (s *QueueService) Read() {}

func (s *QueueService) Create() {}

func (s *QueueService) Update() {}

func (s *QueueService) Delete() {}
