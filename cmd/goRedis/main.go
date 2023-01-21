package main

import (
	"go-redis/api"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	api.RestAPI()
}
