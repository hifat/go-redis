package main

import (
	"go-redis/internal/database"
	"go-redis/internal/model/gormModel"
)

func main() {
	db := database.PostgresDB()
	db.AutoMigrate(&gormModel.Task{})
}
