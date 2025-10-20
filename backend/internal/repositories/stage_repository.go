package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/ent"
)

type StageRepository struct {
	client *ent.Client
}

func NewStageRepository() *StageRepository {
	return &StageRepository{
		client: db.Client,
	}
}

// ([]*ent.Stage, error)
func (s *StageRepository) Read() {}

func (s *StageRepository) Create() {}

func (s *StageRepository) Update() {}

func (s *StageRepository) Delete() {}
