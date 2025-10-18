//go:build wireinject
// +build wireinject

package internal

import (
	"crm.saas/backend/internal/http/handlers"
	"crm.saas/backend/internal/repositories"
	"crm.saas/backend/internal/services"

	"github.com/google/wire"
)

func InitializeChatHandler() *handlers.ChatHandler {
	panic(wire.Build(
		repositories.NewChatRepository,
		services.NewChatService,
		handlers.NewChatHandler,
	))
}

func InitializeCostumerHandler() *handlers.CostumerHandler {
	panic(wire.Build(
		repositories.NewCostumerRepository,
		services.NewCostumerService,
		handlers.NewCostumerHandler,
	))
}
func InitializeDealHandler() *handlers.DealHandler {
	panic(wire.Build(
		repositories.NewDealRepository,
		services.NewDealService,
		handlers.NewDealHandler,
	))
}
func InitializeDepartmentHandler() *handlers.DepartmentHandler {
	panic(wire.Build(
		repositories.NewDepartmentRepository,
		services.NewDepartmentService,
		handlers.NewDepartmentHandler,
	))
}
func InitializeEmployeeHandler() *handlers.EmployeeHandler {
	panic(wire.Build(
		repositories.NewEmployeeRepository,
		services.NewEmployeeService,
		handlers.NewEmployeeHandler,
	))
}
func InitializeEmployeeAuthHandler() *handlers.EmployeeAuthHandler {
	panic(wire.Build(
		repositories.NewEmployeeAuthRepository,
		services.NewEmployeeAuthService,
		handlers.NewEmployeeAuthHandler,
	))
}
func InitializeFieldHandler() *handlers.FieldHandler {
	panic(wire.Build(
		repositories.NewFieldRepository,
		services.NewFieldService,
		handlers.NewFieldHandler,
	))
}
func InitializeFormHandler() *handlers.FormHandler {
	panic(wire.Build(
		repositories.NewFormRepository,
		services.NewFormService,
		handlers.NewFormHandler,
	))
}
func InitializeMessageHandler() *handlers.MessageHandler {
	panic(wire.Build(
		repositories.NewMessageRepository,
		services.NewMessageService,
		handlers.NewMessageHandler,
	))
}
func InitializePipelineHandler() *handlers.PipelineHandler {
	panic(wire.Build(
		repositories.NewPipelineRepository,
		services.NewPipelineService,
		handlers.NewPipelineHandler,
	))
}
func InitializeQueueHandler() *handlers.QueueHandler {
	panic(wire.Build(
		repositories.NewQueueRepository,
		services.NewQueueService,
		handlers.NewQueueHandler,
	))
}
func InitializeStageHandler() *handlers.StageHandler {
	panic(wire.Build(
		repositories.NewStageRepository,
		services.NewStageService,
		handlers.NewStageHandler,
	))
}
