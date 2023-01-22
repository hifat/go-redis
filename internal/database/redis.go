package database

import (
	"os"

	"github.com/go-redis/redis"
	_ "github.com/joho/godotenv/autoload"
)

func Redis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}
