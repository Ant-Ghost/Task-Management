package routes

import (
	"github.com/gin-gonic/gin"

	"zocket/task_manager/controllers"
)

func RegisterTaskRoutes(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/", controllers.GetAllTasks)
	routeGroup.GET("/:id", controllers.GetTask)
	routeGroup.GET("/:id/subtasks", controllers.GetAllSubTasks)
	routeGroup.GET("/filter", controllers.GetTasksByFilters)
	routeGroup.POST("/", controllers.CreateTask)
	routeGroup.POST("/suggestions", controllers.GetAiSuggestion)
	routeGroup.PATCH("/:id", controllers.UpdateTask)
	routeGroup.DELETE("/:id", controllers.DeleteTask)
}
