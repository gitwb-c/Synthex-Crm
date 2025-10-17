package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type EmployeeAuthService struct {
	client *ent.Client
}

func NewEmployeeAuthService() *EmployeeAuthService {
	return &EmployeeAuthService{
		client: db.Client,
	}
}

// ([]*ent.EmployeeAuth, error)
func (s *EmployeeAuthService) Read(data dto.ReadEmployeeAuthReq) {}

func (s *EmployeeAuthService) Create(data dto.ReadEmployeeAuthReq) {}

func (s *EmployeeAuthService) Update(data dto.ReadEmployeeAuthReq) {}

func (s *EmployeeAuthService) Delete(data dto.ReadEmployeeAuthReq) {}
