package main

import (
	"net/http"
	"zocket/task_manager/database"
	"zocket/task_manager/middlewares"
	"zocket/task_manager/openai"
	"zocket/task_manager/slack"
	"zocket/task_manager/websocket"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"zocket/task_manager/controllers"
	"zocket/task_manager/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system env vars")
	}
	go websocket.GlobalHub.Start()

	r := gin.Default()

	database.ConnectDatabase()

	openai.InitOpenAIClient()
	slack.InitiateSlack()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Connected to Task Manager API")
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuth())
	userGroup := auth.Group("/user")
	taskGroup := auth.Group("/task")

	auth.GET("/ws", websocket.GlobalHub.HandleConnections)

	routes.RegisterUserRoutes(userGroup)
	routes.RegisterTaskRoutes(taskGroup)

	r.Run(":8080")
}
