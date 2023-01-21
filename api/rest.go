package api

import (
	"fmt"
	"os"
	"redigo/internal/database"
	"redigo/internal/handler"
	"redigo/internal/repository"
	"redigo/internal/service"

	"github.com/gin-gonic/gin"
)

func RestAPI() {
	r := gin.Default()
	postgresDB := database.PostgresDB()

	newTaskRepo := repository.NewTaskRepository(postgresDB)
	newTaskService := service.NewTaskService(newTaskRepo)
	newTaskHandler := handler.NewTaskHandler(newTaskService)

	taskRoute := r.Group("/tasks")
	{
		taskRoute.GET("", newTaskHandler.Get)
		taskRoute.POST("", newTaskHandler.Store)
		taskRoute.GET("/:taskID", newTaskHandler.Show)
		taskRoute.PATCH("/:taskID", newTaskHandler.Update)
		taskRoute.PATCH("/:taskID/toggle", newTaskHandler.ToggleDone)
		taskRoute.DELETE("/:taskID", newTaskHandler.Delete)
	}

	r.Run(fmt.Sprintf("%s:%s",
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	))
}
