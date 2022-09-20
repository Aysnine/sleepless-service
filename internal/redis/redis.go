package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func New(url string, db int) (client *redis.Client) {
	opt, err := redis.ParseURL(url)

	if err != nil {
		log.Fatalln("RedisParseOptionError:", err)
		return nil
	}

	opt.DB = db

	client = redis.NewClient(opt)

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalln("RedisConnectError:", err)
		return nil
	}

	return client
}
