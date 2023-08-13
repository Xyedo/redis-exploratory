package consumer

import (
	"context"
	"redis-exploratory/queue"

	"github.com/redis/go-redis/v9"
)

func New(rdb *redis.Client) *Consumer {
	return &Consumer{
		db: rdb,
	}
}

type Consumer struct {
	db *redis.Client
}

func (self *Consumer) Consume(ctx context.Context) ([]string, error) {
	return self.db.BRPop(ctx, 0, queue.Notification).Result()

}
