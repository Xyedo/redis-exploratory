package main

import (
	"context"
	"log"
	"redis-exploratory/queue"
	"redis-exploratory/queue/consumer"
	"time"
)

func main() {
	redisClient, err := queue.NewRedisClient(context.TODO())
	if err != nil {
		panic(err)
	}

	c := consumer.New(redisClient)
	for {
		go consume(c, 1)
		go consume(c, 2)
		time.Sleep(3 * time.Second)
	}

}

func consume(c *consumer.Consumer, clientNumber int) {
	s, err := c.Consume(context.TODO())
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("consumer-%d got message %v\n", clientNumber, s)

}
