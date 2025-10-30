package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gitwb-c/crm.saas/backend/internal/domain/contracts"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
	"github.com/gitwb-c/crm.saas/backend/pkg/auth"
	"github.com/gitwb-c/crm.saas/backend/pkg/jwtpkg"
	"github.com/google/uuid"
)

type EmployeeAuthService struct {
	repository *repositories.EmployeeAuthRepository
}

func NewEmployeeAuthService(repository *repositories.EmployeeAuthRepository) *EmployeeAuthService {
	return &EmployeeAuthService{
		repository: repository,
	}
}
func (s *EmployeeAuthService) Read(ctx context.Context, tenantId uuid.UUID) ([]*ent.EmployeeAuth, error) {
	employeeAuth, err := s.repository.Read(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	return employeeAuth, nil
}

func (s *EmployeeAuthService) ReadID(ctx context.Context, id uuid.UUID) (*ent.EmployeeAuth, error) {
	employee, err := s.repository.ReadID(ctx, id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeAuthService) ValidateLogin(ctx context.Context, email string, password string, tenantId uuid.UUID) (*contracts.NewLogin, int, error) {
	bootstrapSig := viewer.Login
	reqCtx := viewer.NewContext(ctx, viewer.UserViewer{TenantID: tenantId, Signature: &bootstrapSig})
	employeeAuth, err := s.repository.ReadEmail(reqCtx, email, tenantId)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	valid := auth.CheckPasswordHash(password, employeeAuth.Password)
	if !valid {
		return nil, http.StatusUnauthorized, fmt.Errorf("invalid credentials")
	}

	tokenStr, sessionId, err := jwtpkg.GenerateToken()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return &contracts.NewLogin{
		SessionId:    sessionId,
		Jwt:          tokenStr,
		EmployeeId:   employeeAuth.Edges.Employee.ID.String(),
		TenantId:     employeeAuth.Edges.Tenant.ID.String(),
		DepartmentId: employeeAuth.Edges.Employee.Edges.Department.ID.String(),
		Time:         time.Now(),
	}, http.StatusOK, nil
}

func (s *EmployeeAuthService) Create(ctx context.Context, input ent.CreateEmployeeAuthInput, client *ent.Client) (*ent.EmployeeAuth, error) {
	if _, err := auth.ValidateNameLenght(input.Name); err != nil {
		return nil, err
	}
	if _, err := auth.ValidatePasswordLenght(input.Password); err != nil {
		return nil, err
	}
	if _, err := auth.ValidateEmail(input.Email); err != nil {
		return nil, err
	}
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	input.Password = hashedPassword
	employeeAuth, err := s.repository.Create(ctx, input, client)
	if err != nil {
		return nil, err
	}
	return employeeAuth, nil
}
func (s *EmployeeAuthService) Update(ctx context.Context, ids []uuid.UUID, input ent.UpdateEmployeeAuthInput) (int, error) {
	n, err := s.repository.Update(ctx, ids, input)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (s *EmployeeAuthService) Delete(ctx context.Context, ids []uuid.UUID) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
