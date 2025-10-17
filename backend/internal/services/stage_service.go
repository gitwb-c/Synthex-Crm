package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type StageService struct {
	client *ent.Client
}

func NewStageService() *StageService {
	return &StageService{
		client: db.Client,
	}
}

// ([]*ent.Stage, error)
func (s *StageService) Read(data dto.ReadStageReq) {}

func (s *StageService) Create(data dto.ReadStageReq) {}

func (s *StageService) Update(data dto.ReadStageReq) {}

func (s *StageService) Delete(data dto.ReadStageReq) {}
