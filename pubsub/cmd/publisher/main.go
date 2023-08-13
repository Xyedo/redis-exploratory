package main

import (
	"context"
	"log"
	"redis-exploratory/pkg/database"
	"redis-exploratory/pubsub/publisher"
	"time"
)

func main() {
	c, err := database.NewRedisClient(context.TODO())
	if err != nil {
		panic(err)
	}

	p := publisher.New(c)
	for {
		s, err := p.Publish()
		if err != nil {
			log.Printf("publish err: %v\n", err)
			continue
		}
		log.Println("published with id:", s)
		time.Sleep(1 * time.Second)
	}

}
