package routes

import (
	"github.com/gin-gonic/gin"

	"zocket/task_manager/controllers"
)

func RegisterUserRoutes(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/me", controllers.GetMyUser)
	routeGroup.PATCH("/:id", controllers.UpdateUser)
	routeGroup.GET("/:id", controllers.GetUser)
	routeGroup.GET("/", controllers.GetAllUsers)
}
