package main

import (
	"context"
	"log"
	"redis-exploratory/optimisticlocking"
	"redis-exploratory/pkg/database"
	"strconv"
	"time"
)

type clientErrorPair struct {
	clientId string
	err      error
}

func main() {
	c, err := database.NewRedisClient(context.TODO())
	if err != nil {
		panic(err)
	}

	client := 100_000
	log.Printf("concurent user who want to bid is %d\n", client)
	userIds := make([]string, client)
	for i := range userIds {
		userIds[i] = strconv.Itoa(i)
	}

	o := optimisticlocking.New(c)
	errChan := make(chan clientErrorPair, client)
	userIdChan := make(chan string)
	for {
		select {
		case pair, ok := <-errChan:
			if ok {
				log.Printf("user %s trying to bid, but %v\n", pair.clientId, pair.err)
			}
		case userId, ok := <-userIdChan:
			if ok {
				log.Println("current user who get the lock is", userId)
			}
		default:
		}

		for i := range userIds {
			go func(i int) {
				err := o.ChangeCurrentGet(userIds[i], func(s string) error {
					if s != "" {
						userIdChan <- userIds[i]
					}

					return nil
				})
				if err != nil {
					errChan <- clientErrorPair{
						clientId: userIds[i],
						err:      err,
					}
				}

			}(i)
		}
		time.Sleep(1 * time.Second)
	}

}
