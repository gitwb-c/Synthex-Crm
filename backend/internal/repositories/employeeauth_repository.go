package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type EmployeeAuthRepository struct {
	client *ent.Client
}

func NewEmployeeAuthRepository() *EmployeeAuthRepository {
	return &EmployeeAuthRepository{
		client: db.Client,
	}
}

// ([]*ent.EmployeeAuth, error)
func (s *EmployeeAuthRepository) Read(data dto.ReadEmployeeAuthReq) {}

func (s *EmployeeAuthRepository) Create(data dto.CreateEmployeeAuthReq) {}

func (s *EmployeeAuthRepository) Update(data dto.UpdateEmployeeAuthReq) {}

func (s *EmployeeAuthRepository) Delete(data dto.DeleteEmployeeAuthReq) {}
