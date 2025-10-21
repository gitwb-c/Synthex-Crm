package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type ChatRepository struct {
	client *ent.Client
}

func NewChatRepository(client *ent.Client) *ChatRepository {
	return &ChatRepository{
		client: client,
	}
}

// ([]*ent.Chat, error)
func (s *ChatRepository) Read(ctx context.Context) {}

func (s *ChatRepository) Create(ctx context.Context) {}

func (s *ChatRepository) Update(ctx context.Context) {}

func (s *ChatRepository) Delete(ctx context.Context) {}
