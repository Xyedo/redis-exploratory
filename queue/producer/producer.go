package producer

import (
	"context"
	"log"
	"redis-exploratory/queue"

	"github.com/go-faker/faker/v4"
	"github.com/redis/go-redis/v9"
)

func New(rdb *redis.Client) *Producer {
	return &Producer{
		db: rdb,
	}
}

type Producer struct {
	db *redis.Client
}

func (self *Producer) Produce(ctx context.Context) {
	str := "hello " + faker.Name()
	log.Println("producing:", str)
	self.db.LPush(ctx, queue.Notification, str)
}
