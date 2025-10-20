//go:build wireinject
// +build wireinject

package wire

import (
	"crm.saas/backend/internal/repositories"
	"crm.saas/backend/internal/services"

	"github.com/google/wire"
)

func InitializeChatHandler() *services.ChatService {
	panic(wire.Build(
		repositories.NewChatRepository,
		services.NewChatService,
	))
}

func InitializeCostumerHandler() *services.CostumerService {
	panic(wire.Build(
		repositories.NewCostumerRepository,
		services.NewCostumerService,
	))
}
func InitializeDealHandler() *services.DealService {
	panic(wire.Build(
		repositories.NewDealRepository,
		services.NewDealService,
	))
}
func InitializeDepartmentHandler() *services.DepartmentService {
	panic(wire.Build(
		repositories.NewDepartmentRepository,
		services.NewDepartmentService,
	))
}
func InitializeEmployeeHandler() *services.EmployeeService {
	panic(wire.Build(
		repositories.NewEmployeeRepository,
		services.NewEmployeeService,
	))
}
func InitializeEmployeeAuthHandler() *services.EmployeeAuthService {
	panic(wire.Build(
		repositories.NewEmployeeAuthRepository,
		services.NewEmployeeAuthService,
	))
}
func InitializeFieldHandler() *services.FieldService {
	panic(wire.Build(
		repositories.NewFieldRepository,
		services.NewFieldService,
	))
}
func InitializeFormHandler() *services.FormService {
	panic(wire.Build(
		repositories.NewFormRepository,
		services.NewFormService,
	))
}
func InitializeMessageHandler() *services.MessageService {
	panic(wire.Build(
		repositories.NewMessageRepository,
		services.NewMessageService,
	))
}
func InitializePipelineHandler() *services.PipelineService {
	panic(wire.Build(
		repositories.NewPipelineRepository,
		services.NewPipelineService,
	))
}
func InitializeQueueHandler() *services.QueueService {
	panic(wire.Build(
		repositories.NewQueueRepository,
		services.NewQueueService,
	))
}
func InitializeStageHandler() *services.StageService {
	panic(wire.Build(
		repositories.NewStageRepository,
		services.NewStageService,
	))
}
