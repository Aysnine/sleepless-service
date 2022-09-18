package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Aysnine/sleepless-service/internal/bridge"
	"github.com/Aysnine/sleepless-service/internal/room"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var dev bool
var port int
var redisURL string

func init() {
	flag.BoolVar(&dev, "dev", false, "local development mode")
	flag.IntVar(&port, "port", 51339, "server start at port")
	flag.StringVar(&redisURL, "redis", "redis://:6379/0", "redis url eg: redis://<user>:<password>@<host>:<port>/<db_number>")
	flag.Parse()
}

func main() {
	app := fiber.New()
	plaza := room.New()
	rdc := bridge.New(redisURL)

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// * Websocket Members
	app.Get("/ws/public", websocket.New(func(conn *websocket.Conn) {
		member := room.NewWebSocketMember(conn)
		key := plaza.Join(member)

		for {
			if msg, err := member.Receive(); err != nil || member.IsKicked() {
				break
			} else {
				plaza.Broadcast(msg)
			}
		}
		plaza.Leave(key)
	}))

	// * Shared Member
	go func() {
		channelName := "room:plaza"

		pubsub := rdc.Subscribe(context.Background(), channelName)
		defer pubsub.Close()

		member := room.NewRedisMember(rdc, pubsub, channelName)
		key := plaza.Join(member)

		var (
			msg []byte
			err error
		)
		for {
			if msg, err = member.Receive(); err != nil || member.IsKicked() {
				break
			} else {
				plaza.Delivery(msg, key)
			}
		}
		plaza.Leave(key)
		log.Fatalln("RedisError:", err)
	}()

	address := ":" + fmt.Sprint(port)
	if dev {
		address = "localhost:" + fmt.Sprint(port)
	}
	log.Fatal(app.Listen(address))
}
