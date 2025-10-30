package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gitwb-c/crm.saas/backend/internal/domain/contracts"
	"github.com/gitwb-c/crm.saas/backend/internal/domain/services"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/employee"
	"github.com/gitwb-c/crm.saas/backend/internal/viewer"
)

func Bootstrap(
	client *ent.Client,
	companyService *services.CompanyService,
	employeeAuthService *services.EmployeeAuthService,
	employeeService *services.EmployeeService,
	departmentService *services.DepartmentService,
	data contracts.BootstrapDtoRequest,
	ctx context.Context,
) (*contracts.NewSetup, error) {
	bootstrapSig := viewer.Bootstrap
	reqCtx := viewer.NewContext(ctx, viewer.UserViewer{
		Signature: &bootstrapSig,
	})

	tx, err := client.Tx(ctx)
	if err != nil {
		log.Printf("failed to start transaction: %v", err)
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	txClient := tx.Client()

	company, err := companyService.Create(reqCtx, ent.CreateCompanyInput{Name: data.CompanyName}, txClient)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			log.Printf("failed to rollback transaction: %v", rerr)
		}
		log.Printf("failed to create company: %v", err)
		return nil, fmt.Errorf("failed to create company: %w", err)
	}

	reqCtx = viewer.NewContext(reqCtx, viewer.UserViewer{
		TenantID:    company.ID,
		Permissions: nil,
	})

	department, err := departmentService.Create(reqCtx, ent.CreateDepartmentInput{
		Name:     "ADMIN.",
		TenantID: company.ID,
	}, txClient)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			log.Printf("failed to rollback transaction: %v", rerr)
		}
		log.Printf("failed to create department: %v", err)
		return nil, fmt.Errorf("failed to create department: %w", err)
	}

	employeeAuth, err := employeeAuthService.Create(reqCtx, ent.CreateEmployeeAuthInput{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		TenantID: company.ID,
	}, txClient)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			log.Printf("failed to rollback transaction: %v", rerr)
		}
		log.Printf("failed to create employee auth: %v", err)
		return nil, fmt.Errorf("failed to create employee auth: %w", err)
	}

	_, err = employeeService.Create(reqCtx, ent.CreateEmployeeInput{
		Name:             data.EmployeeName,
		EmploymentStatus: employee.EmploymentStatusActive,
		DepartmentID:     department.ID,
		EmployeeAuthID:   employeeAuth.ID,
		TenantID:         company.ID,
	}, txClient)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			log.Printf("failed to rollback transaction: %v", rerr)
		}
		log.Printf("failed to create employee: %v", err)
		return nil, fmt.Errorf("failed to create employee: %w", err)
	}

	if err := tx.Commit(); err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &contracts.NewSetup{Time: time.Now()}, nil
}
