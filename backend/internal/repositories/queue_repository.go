package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/google/uuid"
)

type QueueRepository struct {
	client *ent.Client
}

func NewQueueRepository(client *ent.Client) *QueueRepository {
	return &QueueRepository{
		client: client,
	}
}

func (s *QueueRepository) Read(ctx context.Context) ([]*ent.Queue, error) {
	return s.client.Queue.Query().All(ctx)
}

func (s *QueueRepository) Create(ctx context.Context, input ent.CreateQueueInput) (*ent.Queue, error) {
	return s.client.Queue.Create().SetInput(input).Save(ctx)
}
func (s *QueueRepository) UpdateID(ctx context.Context, id string, input ent.UpdateQueueInput) (*ent.Queue, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Queue.UpdateOneID(uuidId).SetInput(input).Save(ctx)
}
func (s *QueueRepository) Delete(ctx context.Context) {}
