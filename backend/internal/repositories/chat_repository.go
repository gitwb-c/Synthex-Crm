package repositories

import (
	"context"
	"fmt"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/chat"
	"github.com/google/uuid"
)

type ChatRepository struct {
	client *ent.Client
}

func NewChatRepository(client *ent.Client) *ChatRepository {
	return &ChatRepository{
		client: client,
	}
}

func (s *ChatRepository) Read(ctx context.Context) ([]*ent.Chat, error) {
	return s.client.Chat.Query().All(ctx)
}

func (s *ChatRepository) Create(ctx context.Context, input ent.CreateChatInput) (*ent.Chat, error) {
	return s.client.Chat.Create().SetInput(input).Save(ctx)
}
func (s *ChatRepository) UpdateID(ctx context.Context, id string, input ent.UpdateChatInput) (*ent.Chat, error) {
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.client.Chat.UpdateOneID(uuidId).SetInput(input).Save(ctx)
}
func (s *ChatRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	tx, err := s.client.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = tx.Chat.Delete().Where(chat.IDIn(ids...)).Exec(ctx)
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
