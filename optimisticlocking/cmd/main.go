package main

import (
	"context"
	"log"
	"redis-exploratory/optimisticlocking"
	"redis-exploratory/pkg/database"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

func main() {
	c, err := database.NewRedisClient(context.TODO())
	if err != nil {
		panic(err)
	}

	client := 4
	userIds := make([]string, client)
	for i := range userIds {
		userIds[i] = faker.Name(options.WithRandomMapAndSliceMinSize(uint(client)))
	}

	o := optimisticlocking.New(c)
	errChan := make(chan error)
	userIdChan := make(chan string)
	for {
		select {
		case err, ok := <-errChan:
			if ok {
				log.Println(err)
			}
		case userId, ok := <-userIdChan:
			if ok {
				log.Println("current user who get the lock is", userId)
			}
		default:
		}

		for i := range userIds {
			go func(i int) {
				userId, err := o.ChangeCurrentGet(userIds[i])
				if err != nil {
					errChan <- err
				}
				if userId != "" {
					userIdChan <- userId
				}
			}(i)
		}
		time.Sleep(1 * time.Second)
	}

}
