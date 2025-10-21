package services

import (
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
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
func (s *StageService) Read() {}

func (s *StageService) Create() {}

func (s *StageService) Update() {}

func (s *StageService) Delete() {}
