package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type MessageRepository struct {
	client *ent.Client
}

func NewMessageRepository(client *ent.Client) *MessageRepository {
	return &MessageRepository{
		client: client,
	}
}

// ([]*ent.Message, error)
func (s *MessageRepository) Read(ctx context.Context) {}

func (s *MessageRepository) Create(ctx context.Context) {}

func (s *MessageRepository) Update(ctx context.Context) {}

func (s *MessageRepository) Delete(ctx context.Context) {}
