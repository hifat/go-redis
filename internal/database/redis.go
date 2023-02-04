package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
)

func Redis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}
