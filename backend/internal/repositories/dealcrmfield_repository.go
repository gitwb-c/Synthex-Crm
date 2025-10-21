package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type FormRepository struct {
	client *ent.Client
}

func NewFormRepository(client *ent.Client) *FormRepository {
	return &FormRepository{
		client: client,
	}
}

// ([]*ent.Form, error)
func (s *FormRepository) Read(ctx context.Context) {}

func (s *FormRepository) Create(ctx context.Context) {}

func (s *FormRepository) Update(ctx context.Context) {}

func (s *FormRepository) Delete(ctx context.Context) {}
