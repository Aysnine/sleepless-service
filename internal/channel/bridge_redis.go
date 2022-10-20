package channel

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisBridge struct {
	ChannelName string
	Pubsub      *redis.PubSub
	Client      *redis.Client
}

func (m *RedisBridge) Receive() (msg []byte, err error) {
	var x *redis.Message = <-m.Pubsub.Channel()
	return []byte(x.Payload), nil
}

func (m *RedisBridge) Reply(msg []byte) {
	if err := m.Client.Publish(context.Background(), m.ChannelName, msg).Err(); err != nil {
		log.Println("RedisPublishError:", err.Error(), m.ChannelName)
	}
}

func NewRedisBridge(client *redis.Client, pubsub *redis.PubSub, channelName string) *RedisBridge {
	bridge := &RedisBridge{
		ChannelName: channelName,
		Pubsub:      pubsub,
		Client:      client,
	}

	return bridge
}
