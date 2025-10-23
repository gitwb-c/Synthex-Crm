package services

import (
	"context"
	"fmt"
	"time"

	"github.com/gitwb-c/crm.saas/backend/internal/domain/contracts"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/gitwb-c/crm.saas/backend/pkg/auth"
	"github.com/gitwb-c/crm.saas/backend/pkg/jwtpkg"
)

type EmployeeAuthService struct {
	repository *repositories.EmployeeAuthRepository
}

func NewEmployeeAuthService(repository *repositories.EmployeeAuthRepository) *EmployeeAuthService {
	return &EmployeeAuthService{
		repository: repository,
	}
}
func (s *EmployeeAuthService) Read(ctx context.Context) ([]*ent.EmployeeAuth, error) {
	employeeAuth, err := s.repository.Read(ctx)
	if err != nil {
		return nil, err
	}
	return employeeAuth, nil
}

func (s *EmployeeAuthService) ValidateLogin(ctx context.Context, email string, password string) (contracts.NewLogin, error) {
	employeeAuth, err := s.repository.ReadEmail(ctx, email)
	if err != nil {
		return contracts.NewLogin{}, err
	}
	valid := auth.CheckPasswordHash(password, employeeAuth.Password)
	if !valid {
		return contracts.NewLogin{}, fmt.Errorf("invalid credentials")
	}
	tokenStr, err := jwtpkg.GenerateToken(
		employeeAuth.Edges.Employee.ID.String(),
		employeeAuth.Edges.Employee.Edges.Company.ID.String(),
		employeeAuth.Edges.Employee.Edges.Department.ID.String(),
	)
	if err != nil {
		return contracts.NewLogin{}, err
	}
	return contracts.NewLogin{
		Jwt:  tokenStr,
		Time: time.Now(),
	}, nil
}

func (s *EmployeeAuthService) Create(ctx context.Context, input ent.CreateEmployeeAuthInput) (*ent.EmployeeAuth, error) {
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
	employeeAuth, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return employeeAuth, nil
}
func (s *EmployeeAuthService) UpdateID(ctx context.Context, id string, input ent.UpdateEmployeeAuthInput) (*ent.EmployeeAuth, error) {
	employeeAuth, err := s.repository.UpdateID(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return employeeAuth, nil
}

func (s *EmployeeAuthService) DeleteID(ctx context.Context, id string) error {
	err := s.repository.DeleteID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
