package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Nil = redis.Nil

func New(url string, db int) (client *redis.Client, err error) {
	opt, err := redis.ParseURL(url)

	if err != nil {
		return nil, err
	}

	opt.DB = db

	client = redis.NewClient(opt)

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
