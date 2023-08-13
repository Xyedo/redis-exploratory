package database

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	s, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	if s != "PONG" {
		return nil, errors.New("invalid redis response")
	}

	return rdb, nil
}
