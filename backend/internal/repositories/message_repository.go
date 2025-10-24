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

func (s *MessageRepository) Read(ctx context.Context) ([]*ent.Message, error) {
	return s.client.Message.Query().All(ctx)
}

func (s *MessageRepository) Create(ctx context.Context, input ent.CreateMessageInput) (*ent.Message, error) {
	return s.client.Message.Create().SetInput(input).Save(ctx)
}
func (s *MessageRepository) UpdateID(ctx context.Context, id string, input ent.UpdateMessageInput) (*ent.Message, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Message.UpdateOneID(uuidId).SetInput(input).Save(ctx)
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
