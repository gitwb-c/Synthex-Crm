package services

import (
	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/dto"
	"crm.saas/backend/internal/ent"
)

type DepartmentService struct {
	client *ent.Client
}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{
		client: db.Client,
	}
}

// ([]*ent.Department, error)
func (s *DepartmentService) Read(data dto.ReadDepartmentReq) {}

func (s *DepartmentService) Create(data dto.ReadDepartmentReq) {}

func (s *DepartmentService) Update(data dto.ReadDepartmentReq) {}

func (s *DepartmentService) Delete(data dto.ReadDepartmentReq) {}
