package optimisticlocking

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrAlreadyTaken = errors.New("key is already taken")
var ErrExpiredLock = errors.New("key is already expired")

func New(rdb *redis.Client) *Optimistic {
	return &Optimistic{
		db: rdb,
	}
}

type Optimistic struct {
	db *redis.Client
}

func (self Optimistic) ChangeCurrentGet(userId string, cb func(s string) error) error {
	ctx := context.Background()

	ok, err := self.db.SetNX(ctx, AuctionUser, userId, 10*time.Second).Result()
	if err != nil {
		return err
	}

	if !ok {
		return ErrAlreadyTaken
	}

	err = cb(userId)
	if err != nil {
		if errRel := self.release(ctx, AuctionUser); errRel != nil {
			return fmt.Errorf("from err %w: unable to release %w", err, errRel)
		}

		return err
	}

	return self.release(ctx, AuctionUser)
}

func (self Optimistic) release(ctx context.Context, key string) error {
	status, err := self.db.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	if status == -1 {
		return ErrExpiredLock
	}

	return nil
}
