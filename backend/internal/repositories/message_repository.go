package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
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
func (s *MessageRepository) DeleteID(ctx context.Context, id string) error {
	uuidId, e := uuid.Parse(id)
	if e != nil {
		return e
	}
	err := s.client.Message.DeleteOneID(uuidId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
