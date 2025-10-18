package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type QueueRepository struct {
	client *ent.Client
}

func NewQueueRepository() *QueueRepository {
	return &QueueRepository{
		client: db.Client,
	}
}

// ([]*ent.Queue, error)
func (s *QueueRepository) Read(data dto.ReadQueueReq) {}

func (s *QueueRepository) Create(data dto.CreateQueueReq) {}

func (s *QueueRepository) Update(data dto.UpdateQueueReq) {}

func (s *QueueRepository) Delete(data dto.DeleteQueueReq) {}
