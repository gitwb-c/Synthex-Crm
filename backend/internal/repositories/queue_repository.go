package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type QueueRepository struct {
	client *ent.Client
}

func NewQueueRepository(client *ent.Client) *QueueRepository {
	return &QueueRepository{
		client: client,
	}
}

// ([]*ent.Queue, error)
func (s *QueueRepository) Read(ctx context.Context) {}

func (s *QueueRepository) Create(ctx context.Context) {}

func (s *QueueRepository) Update(ctx context.Context) {}

func (s *QueueRepository) Delete(ctx context.Context) {}
