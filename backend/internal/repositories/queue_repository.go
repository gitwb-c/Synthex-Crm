package repositories

import (
	"crm.saas/backend/internal/db"
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
func (s *QueueRepository) Read() {}

func (s *QueueRepository) Create() {}

func (s *QueueRepository) Update() {}

func (s *QueueRepository) Delete() {}
