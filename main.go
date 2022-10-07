package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Aysnine/sleepless-service/internal/channel"
	"github.com/Aysnine/sleepless-service/internal/platform/redis"
	"github.com/Aysnine/sleepless-service/internal/platform/wechat"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
)

var dev bool
var port int
var redisURL string
var wxAppId string
var wxAppSecret string
var jwtSecret string

func init() {
	flag.BoolVar(&dev, "dev", false, "local development mode")
	flag.IntVar(&port, "port", 51339, "server start at port")

	defaultRedisURL := os.Getenv("REDIS_URL")
	if len(defaultRedisURL) == 0 {
		defaultRedisURL = "redis://:6379/0"
	}
	flag.StringVar(&redisURL, "redis", defaultRedisURL, "redis url eg: redis://<user>:<password>@<host>:<port>/<db_number>. or REDIS_URL environment variable")

	defaultJwtSecret := os.Getenv("JWT_SECRET")
	if len(defaultJwtSecret) == 0 {
		defaultJwtSecret = "SECRET"
	}
	flag.StringVar(&jwtSecret, "jwt-secret", defaultJwtSecret, "JWT signing key")

	defaultWxAppId := os.Getenv("WX_APP_ID")
	if len(defaultWxAppId) == 0 {
		defaultWxAppId = "WX_APP_ID"
	}
	flag.StringVar(&wxAppId, "wx-app-id", defaultWxAppId, "wechat mini program app id")

	defaultWxAppSecret := os.Getenv("WX_APP_SECRET")
	if len(defaultWxAppSecret) == 0 {
		defaultWxAppSecret = "WX_APP_SECRET"
	}
	flag.StringVar(&wxAppSecret, "wx-app-secret", defaultWxAppSecret, "wechat mini program app secret")

	flag.Parse()
}

func main() {
	app := fiber.New()
	plaza := channel.New()
	rdc, err := redis.New(redisURL, 0)
	validate := validator.New()
	wechatClient := wechat.New(wxAppId, wxAppSecret)

	if err != nil {
		log.Fatalln("NewRedisError", err)
	}

	app.Post("/funny/wechat-mini-program-login", func(c *fiber.Ctx) error {
		body := &struct {
			Code string `json:"code" validate:"required"`
		}{}

		if err := c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if err := validate.Struct(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		resp, err := wechatClient.JsCodeToSession(body.Code)
		if err != nil || resp.ErrCode != 0 {
			var message string
			if err != nil {
				message = err.Error()
			}
			if resp.ErrCode != 0 {
				message = "Third party service error"
				fmt.Printf("wechatClient.JsCodeToSession(%s) -> %d: %s", body.Code, resp.ErrCode, resp.ErrMsg)
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": message,
			})
		}

		// Create token
		expireInterval := time.Hour * 24 * 7
		if dev {
			expireInterval = time.Minute * 3
		}
		claims := jwt.MapClaims{
			"from":       "WxApp",
			"openId":     resp.OpenId,
			"sessionKey": resp.SessionKey,
			"unionId":    resp.UnionId,

			// ! special key 'exp'
			"exp": time.Now().Add(expireInterval).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{"token": t})
	})

	app.Use(jwtWare.New(jwtWare.Config{
		SigningKey: []byte(jwtSecret),
	}))

	app.Get("/funny/test", func(c *fiber.Ctx) error {
		return c.SendString("hello!")
	})

	app.Use("/funny/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// * Websocket Members
	app.Get("/funny/ws/plaza", websocket.New(func(conn *websocket.Conn) {
		member := channel.NewWebSocketMember(conn)
		key := plaza.Join(member)

		for {
			if msg, err := member.Receive(); err != nil || member.IsKicked() {
				break
			} else {
				go plaza.Broadcast(msg)
			}
		}

		plaza.Leave(key)
	}))

	// * Redis Bridge
	go func() {
		channelName := "plaza"

		pubsub := rdc.Subscribe(context.Background(), channelName)
		defer pubsub.Close()

		bridge := channel.NewRedisBridge(rdc, pubsub, channelName)
		plaza.SetBridge(bridge)

		var (
			msg []byte
			err error
		)
		for {
			if msg, err = bridge.Receive(); err != nil {
				break
			} else {
				go plaza.Delivery(msg)
			}
		}

		plaza.RemoveBridge()
		log.Fatalln("RedisBridgeError:", err)
	}()

	address := ":" + fmt.Sprint(port)
	if dev {
		address = "localhost:" + fmt.Sprint(port)
	}
	log.Fatal(app.Listen(address))
}
