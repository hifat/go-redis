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
	rdb := database.Redis()

	newTaskRepo := repository.NewTaskRepository(postgresDB)
	newTaskService := service.NewTaskService(newTaskRepo, rdb)
	newTaskHandler := handler.NewTaskHandler(newTaskService)

	r.GET("/ht", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

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
