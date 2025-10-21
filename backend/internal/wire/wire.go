//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/repositories"
	"github.com/gitwb-c/crm.saas/backend/internal/services"
	"github.com/google/wire"
)

func InitializeChatService(client *ent.Client) *services.ChatService {
	panic(wire.Build(
		services.NewChatService,
		InitializeChatRepository,
	))
}

func InitializeCostumerService(client *ent.Client) *services.CostumerService {
	panic(wire.Build(
		services.NewCostumerService,
		InitializeCostumerRepository,
	))
}

func InitializeDealService(client *ent.Client) *services.DealService {
	panic(wire.Build(
		services.NewDealService,
		InitializeDealRepository,
	))
}

func InitializeDepartmentService(client *ent.Client) *services.DepartmentService {
	panic(wire.Build(
		services.NewDepartmentService,
		InitializeDepartmentRepository,
	))
}

func InitializeEmployeeService(client *ent.Client) *services.EmployeeService {
	panic(wire.Build(
		services.NewEmployeeService,
		InitializeEmployeeRepository,
	))
}

func InitializeEmployeeAuthService(client *ent.Client) *services.EmployeeAuthService {
	panic(wire.Build(
		services.NewEmployeeAuthService,
		InitializeEmployeeAuthRepository,
	))
}

func InitializeFieldService(client *ent.Client) *services.FieldService {
	panic(wire.Build(
		services.NewFieldService,
		InitializeFieldRepository,
	))
}

func InitializeFormService(client *ent.Client) *services.FormService {
	panic(wire.Build(
		services.NewFormService,
		InitializeFormRepository,
	))
}

func InitializeMessageService(client *ent.Client) *services.MessageService {
	panic(wire.Build(
		services.NewMessageService,
		InitializeMessageRepository,
	))
}

func InitializePipelineService(client *ent.Client) *services.PipelineService {
	panic(wire.Build(
		services.NewPipelineService,
		InitializePipelineRepository,
	))
}

func InitializeQueueService(client *ent.Client) *services.QueueService {
	panic(wire.Build(
		services.NewQueueService,
		InitializeQueueRepository,
	))
}

func InitializeStageService(client *ent.Client) *services.StageService {
	panic(wire.Build(
		services.NewStageService,
		InitializeStageRepository,
	))
}

func InitializeChatRepository(client *ent.Client) *repositories.ChatRepository {
	panic(wire.Build(repositories.NewChatRepository))
}

func InitializeCostumerRepository(client *ent.Client) *repositories.CostumerRepository {
	panic(wire.Build(repositories.NewCostumerRepository))
}

func InitializeDealRepository(client *ent.Client) *repositories.DealRepository {
	panic(wire.Build(repositories.NewDealRepository))
}

func InitializeDepartmentRepository(client *ent.Client) *repositories.DepartmentRepository {
	panic(wire.Build(repositories.NewDepartmentRepository))
}

func InitializeEmployeeRepository(client *ent.Client) *repositories.EmployeeRepository {
	panic(wire.Build(repositories.NewEmployeeRepository))
}

func InitializeEmployeeAuthRepository(client *ent.Client) *repositories.EmployeeAuthRepository {
	panic(wire.Build(repositories.NewEmployeeAuthRepository))
}

func InitializeFieldRepository(client *ent.Client) *repositories.FieldRepository {
	panic(wire.Build(repositories.NewFieldRepository))
}

func InitializeFormRepository(client *ent.Client) *repositories.FormRepository {
	panic(wire.Build(repositories.NewFormRepository))
}

func InitializeMessageRepository(client *ent.Client) *repositories.MessageRepository {
	panic(wire.Build(repositories.NewMessageRepository))
}

func InitializePipelineRepository(client *ent.Client) *repositories.PipelineRepository {
	panic(wire.Build(repositories.NewPipelineRepository))
}

func InitializeQueueRepository(client *ent.Client) *repositories.QueueRepository {
	panic(wire.Build(repositories.NewQueueRepository))
}

func InitializeStageRepository(client *ent.Client) *repositories.StageRepository {
	panic(wire.Build(repositories.NewStageRepository))
}
