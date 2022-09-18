package room

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisMember struct {
	Kicked      bool
	ChannelName string
	Pubsub      *redis.PubSub
	Client      *redis.Client
}

func (m *RedisMember) Receive() (msg []byte, err error) {
	var x *redis.Message = <-m.Pubsub.Channel()
	return []byte(x.Payload), nil
}

func (m *RedisMember) Reply(msg []byte) {
	if err := m.Client.Publish(context.Background(), m.ChannelName, msg).Err(); err != nil {
		log.Println("RedisPublishError:", err, m.ChannelName)
	}
}

func (m *RedisMember) Kick() {
	m.Kicked = true
}

func (m *RedisMember) IsKicked() bool {
	return m.Kicked
}

func NewRedisMember(client *redis.Client, pubsub *redis.PubSub, channelName string) *RedisMember {
	member := &RedisMember{
		Kicked:      false,
		ChannelName: channelName,
		Pubsub:      pubsub,
		Client:      client,
	}

	return member
}
