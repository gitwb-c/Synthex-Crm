package services

import (
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/repositories"
)

type StageService struct {
	repository *repositories.StageRepository
}

func NewStageService(repository *repositories.StageRepository) *StageService {
	return &StageService{
		repository: repository,
	}
}

// ([]*ent.Stage, error)
func (s *StageService) Read(data dto.ReadStageReq) {}

func (s *StageService) Create(data dto.CreateStageReq) {}

func (s *StageService) Update(data dto.UpdateStageReq) {}

func (s *StageService) Delete(data dto.DeleteStageReq) {}
