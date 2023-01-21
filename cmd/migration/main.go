package main

import (
	"redigo/internal/database"
	"redigo/internal/model/gormModel"
)

func main() {
	db := database.PostgresDB()
	db.AutoMigrate(&gormModel.Task{})
}
