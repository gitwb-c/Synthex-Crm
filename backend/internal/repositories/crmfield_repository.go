package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type FieldRepository struct {
	client *ent.Client
}

func NewFieldRepository(client *ent.Client) *FieldRepository {
	return &FieldRepository{
		client: client,
	}
}

// ([]*ent.Field, error)
func (s *FieldRepository) Read(ctx context.Context) {}

func (s *FieldRepository) Create(ctx context.Context) {}

func (s *FieldRepository) Update(ctx context.Context) {}

func (s *FieldRepository) Delete(ctx context.Context) {}
