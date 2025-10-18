package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
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
func (s *StageRepository) Read(data dto.ReadStageReq) {}

func (s *StageRepository) Create(data dto.CreateStageReq) {}

func (s *StageRepository) Update(data dto.UpdateStageReq) {}

func (s *StageRepository) Delete(data dto.DeleteStageReq) {}
