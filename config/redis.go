package config

import (
	"context"

	"golang-api/logger"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

var Ctx = context.Background()

func ConnectRedis() {

	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := RDB.Ping(Ctx).Result()

	if err != nil {

		logger.ErrorLogger.Println(
			"Redis connection failed:",
			err,
		)

		panic(err)
	}

	logger.InfoLogger.Println(
		"Redis connected successfully",
	)
}
