package repositories

import (
	"crm.saas/backend/internal/db"
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
func (s *EmployeeAuthRepository) Read() {}

func (s *EmployeeAuthRepository) Create() {}

func (s *EmployeeAuthRepository) Update() {}

func (s *EmployeeAuthRepository) Delete() {}
