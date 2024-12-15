package main

import (
	"context"
	"log"
	"redis-exploratory/optimisticlocking"
	"redis-exploratory/pkg/database"
	"strconv"
	"time"
)

func main() {
	c, err := database.NewRedisClient(context.TODO())
	if err != nil {
		panic(err)
	}

	client := 100
	userIds := make([]string, client)
	for i := range userIds {
		userIds[i] = strconv.Itoa(i)
	}

	o := optimisticlocking.New(c)
	errChan := make(chan error, client)
	userIdChan := make(chan string, client)
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
				log.Printf("user-%d wanna change the bid\n", i)
				err := o.ChangeCurrentGet(userIds[i], func(s string) error {
					if s != "" {
						userIdChan <- userIds[i]
					}

					return nil
				})
				if err != nil {
					errChan <- err
				}

			}(i)
		}
		time.Sleep(1 * time.Second)
	}

}
