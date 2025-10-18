package repositories

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type DepartmentRepository struct {
	client *ent.Client
}

func NewDepartmentRepository() *DepartmentRepository {
	return &DepartmentRepository{
		client: db.Client,
	}
}

// ([]*ent.Department, error)
func (s *DepartmentRepository) Read(data dto.ReadDepartmentReq) {}

func (s *DepartmentRepository) Create(data dto.CreateDepartmentReq) {}

func (s *DepartmentRepository) Update(data dto.UpdateDepartmentReq) {}

func (s *DepartmentRepository) Delete(data dto.DeleteDepartmentReq) {}
