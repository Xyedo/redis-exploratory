package optimisticlocking

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

func New(rdb *redis.Client) *Optimistic {
	return &Optimistic{
		db: rdb,
	}
}

type Optimistic struct {
	db *redis.Client
}

func (self Optimistic) ChangeCurrentGet(userId string) (string, error) {
	ctx := context.Background()
	err := self.db.Watch(ctx, func(tx *redis.Tx) error {
		s, err := tx.Set(ctx, AuctionUser, userId, 0).Result()
		if err != nil {
			return err
		}
		if s != "OK" {
			return errors.New("idk what happened, but returned is " + s)
		}

		return nil

	}, AuctionUser)
	if err != nil {
		return "", err
	}
	if err != nil && errors.Is(err, redis.Nil) {
		return "", nil
	}

	return userId, nil

}
