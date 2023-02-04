package main

import (
	"log"
	"redigo/internal/database"
	"redigo/internal/repository"
	"redigo/internal/service"
)

func main() {
	newTaskRepository := repository.NewTaskRepository(database.PostgresDB())
	newTaskService := service.NewTaskService(newTaskRepository, nil)

	err := newTaskService.MockData(5000)
	if err != nil {
		log.Println(err.Error())
	}
}
