package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/queue"
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

func (s *QueueRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Queue, error) {
	return s.client.Queue.Query().Where(queue.TenantIdEQ(tenantId)).All(ctx)
}

func (s *QueueRepository) Create(ctx context.Context, input ent.CreateQueueInput) (*ent.Queue, error) {
	return s.client.Queue.Create().SetInput(input).Save(ctx)
}
func (s *QueueRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateQueueInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Queue.Update().Where(queue.IDIn(ids...)).SetInput(input).Save(ctx)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	return n, nil
}
func (s *QueueRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Queue.Delete().Where(queue.IDIn(ids...)).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
