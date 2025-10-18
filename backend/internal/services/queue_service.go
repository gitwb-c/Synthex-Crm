package services

import (
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/repositories"
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
func (s *QueueService) Read(data dto.ReadQueueReq) {}

func (s *QueueService) Create(data dto.CreateQueueReq) {}

func (s *QueueService) Update(data dto.UpdateQueueReq) {}

func (s *QueueService) Delete(data dto.DeleteQueueReq) {}
