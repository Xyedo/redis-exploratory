package publisher

import (
	"context"
	"redis-exploratory/pubsub"

	"github.com/go-faker/faker/v4"
	"github.com/redis/go-redis/v9"
)

func New(rdb *redis.Client) *Publisher {
	return &Publisher{
		db: rdb,
	}
}

type Publisher struct {
	db *redis.Client
}

func (self Publisher) Publish() (string, error) {
	s := faker.Name()
	faker.Sentence()
	return self.db.XAdd(context.TODO(), &redis.XAddArgs{
		Stream: pubsub.UserChange,
		Values: map[string]any{
			"username": "xyedo",
			"name":     s,
		},
	}).Result()
}
