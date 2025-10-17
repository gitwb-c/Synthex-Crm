package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	pipelineRouter(r)
	employeeAuthRouter(r)
	employeeRouter(r)
	chatRouter(r)
	dealRouter(r)
	departmentRouter(r)
	messageRouter(r)
	queueRouter(r)
	stageRouter(r)
	formRouter(r)
	costumerRouter(r)
	fieldRouter(r)
}
