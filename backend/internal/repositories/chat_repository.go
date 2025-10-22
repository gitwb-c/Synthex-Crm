package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
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
func (s *ChatRepository) DeleteID(ctx context.Context, id string) error {
	uuidId, e := uuid.Parse(id)
	if e != nil {
		return e
	}
	err := s.client.Chat.DeleteOneID(uuidId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
