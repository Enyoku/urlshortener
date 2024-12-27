package storage

import (
	"servcached/pkg/config"

	"github.com/redis/go-redis/v9"
)

func New(config *config.APIConfig) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	return redisClient, nil
}
