package repositories

import (
	"context"

	"github.com/gitwb-c/crm.saas/backend/internal/ent"
)

type StageRepository struct {
	client *ent.Client
}

func NewStageRepository(client *ent.Client) *StageRepository {
	return &StageRepository{
		client: client,
	}
}

// ([]*ent.Stage, error)
func (s *StageRepository) Read(ctx context.Context) {}

func (s *StageRepository) Create(ctx context.Context) {}

func (s *StageRepository) Update(ctx context.Context) {}

func (s *StageRepository) Delete(ctx context.Context) {}
