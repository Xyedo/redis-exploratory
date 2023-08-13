package main

import (
	"context"
	"redis-exploratory/pkg/database"
	"redis-exploratory/queue/producer"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := database.NewRedisClient(ctx)
	if err != nil {
		panic(err)
	}

	p := producer.New(c)
	for {
		produceRoutine(p)
	}
}

func produceRoutine(p *producer.Producer) {
	p.Produce(context.TODO())
	time.Sleep(1 * time.Second)
}
