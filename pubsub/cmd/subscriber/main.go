package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"redis-exploratory/pubsub"
	"redis-exploratory/pubsub/subscriber"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	c, err := pubsub.NewRedisClient(context.TODO())
	if err != nil {
		panic(err)
	}

	p := subscriber.New(c)
	go publisher(1, p)
	go publisher(2, p)
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
}

func publisher(clientNumber int, p *subscriber.Subscriber) {
	var lastId string
	for {
		s, err := p.Subscribe(clientNumber, lastId)
		if err != nil && !errors.Is(err, redis.Nil) {
			log.Printf("subscriber-%d err: %v\n", clientNumber, err)
			continue
		}
		if s != "" {
			lastId = s
		}
		time.Sleep(1 * time.Second)
	}
}
