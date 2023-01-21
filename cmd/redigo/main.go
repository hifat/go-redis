package main

import (
	"redigo/api"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	api.RestAPI()
}
