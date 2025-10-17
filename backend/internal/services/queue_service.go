package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type QueueService struct {
	client *ent.Client
}

func NewQueueService() *QueueService {
	return &QueueService{
		client: db.Client,
	}
}

// ([]*ent.Queue, error)
func (s *QueueService) Read(data dto.ReadQueueReq) {}

func (s *QueueService) Create(data dto.ReadQueueReq) {}

func (s *QueueService) Update(data dto.ReadQueueReq) {}

func (s *QueueService) Delete(data dto.ReadQueueReq) {}
