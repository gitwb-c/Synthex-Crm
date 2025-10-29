package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/message"
	"github.com/google/uuid"
)

type MessageRepository struct {
	client *ent.Client
}

func NewMessageRepository(client *ent.Client) *MessageRepository {
	return &MessageRepository{
		client: client,
	}
}

func (s *MessageRepository) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.Message, error) {
	return s.client.Message.Query().Where(message.TenantIdEQ(tenantId)).All(ctx)
}

func (s *MessageRepository) Create(ctx context.Context, input ent.CreateMessageInput) (*ent.Message, error) {
	return s.client.Message.Create().SetInput(input).Save(ctx)
}
func (s *MessageRepository) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateMessageInput) (int, error) {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	n, err := tx.Message.Update().Where(message.IDIn(ids...)).SetInput(input).Save(ctx)
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
func (s *MessageRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Message.Delete().Where(message.IDIn(ids...)).Exec(ctx)
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
