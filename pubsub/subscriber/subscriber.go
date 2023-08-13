package subscriber

import (
	"context"
	"log"
	"redis-exploratory/pubsub"

	"github.com/redis/go-redis/v9"
)

func New(rdb *redis.Client) *Subscriber {
	return &Subscriber{
		db: rdb,
	}
}

type Subscriber struct {
	db *redis.Client
}

func (self *Subscriber) Subscribe(clientNumber int, lastId string) (string, error) {
	if lastId == "" {
		lastId = "0-0"
	}
	streams, err := self.db.XReadStreams(context.TODO(), pubsub.UserChange, lastId).Result()
	if err != nil {
		return "", err
	}

	for _, stream := range streams {
		for _, message := range stream.Messages {
			log.Printf("subscriber-%d got this: %v\n", clientNumber, message.Values)
		}
	}

	if len(streams) == 0 || len(streams[0].Messages) == 0 {
		return "", nil
	}

	return streams[0].Messages[len(streams[0].Messages)-1].ID, nil

}
