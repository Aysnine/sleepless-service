package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Aysnine/sleepless-service/internal/channel"
	"github.com/Aysnine/sleepless-service/internal/message"
	"github.com/Aysnine/sleepless-service/internal/platform/redis"
	"github.com/Aysnine/sleepless-service/internal/platform/wechat"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/proto"
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
		log.Fatalln("NewRedisError", err.Error())
	}

	app.Static("/", "./public")

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

			// ! common key
			"memberId": resp.OpenId,

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

	app.Use("/funny", jwtWare.New(jwtWare.Config{
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

	plazaKey := "room:plaza"
	plazaTidCounterKey := plazaKey + ":tid-counter"
	plazaTidStart := "0"

	// TODO support expiration when service instance to zero
	if err := rdc.SetNX(context.Background(), plazaTidCounterKey, plazaTidStart, 0).Err(); err != nil {
		log.Fatalln("RedisSetNXError", err.Error())
	}

	// * Websocket Members
	app.Get("/funny/ws/plaza", websocket.New(func(conn *websocket.Conn) {
		user := conn.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		memberId := claims["memberId"].(string)

		memberKey := plazaKey + ":member:" + memberId
		tidFieldName := "tid"

		tid := int32(-1)

		if existedTid, err := rdc.HGet(context.Background(), memberKey, tidFieldName).Result(); err != nil && err != redis.Nil {
			fmt.Println("RedisHGetError", err.Error())
		} else {
			if len(existedTid) == 0 {
				if newTid, err := rdc.Incr(context.Background(), plazaTidCounterKey).Result(); err != nil {
					fmt.Println("RedisIncrError", err.Error())
				} else {
					if err := rdc.HSetNX(context.Background(), memberKey, tidFieldName, newTid).Err(); err != nil {
						fmt.Println("RedisHSetNXError", err.Error())
					} else {
						tid = int32(newTid)
					}
				}
			} else {
				if i64, err := strconv.ParseInt(existedTid, 10, 32); err != nil {
					fmt.Println("ParseIntError", existedTid, err.Error())
				} else {
					tid = int32(i64)
				}
			}
		}

		if tid == -1 {
			return
		}

		member := channel.NewWebSocketMember(conn)
		key := plaza.Join(member)

		if msgUnderlay, err := proto.Marshal(
			&message.PublicMessage{
				Action: &message.PublicMessage_Underlay_{
					Underlay: &message.PublicMessage_Underlay{
						Tid: tid,
					},
				},
			},
		); err != nil {
			// TODO ignore log
			fmt.Println("ProtoBufferMarshalError", err.Error())
			return
		} else {
			go member.Reply(msgUnderlay)
		}

		if msgJoin, err := proto.Marshal(
			&message.PublicMessage{
				Action: &message.PublicMessage_Join_{
					Join: &message.PublicMessage_Join{
						Tid: tid,
					},
				},
			},
		); err != nil {
			// TODO ignore log
			fmt.Println("ProtoBufferMarshalError", err.Error())
			return
		} else {
			go plaza.Broadcast(msgJoin)
		}

		exceptionCounter := 0

		for {
			if member.IsKicked() || exceptionCounter > 64 {
				break
			}

			if msg, err := member.Receive(); err != nil {
				exceptionCounter += 1

				// fmt.Println("WsMemberReceiveError", err.Error())
			} else {
				upcomingMsg := message.UpcomingMessage{}

				if err := proto.Unmarshal(msg, &upcomingMsg); err != nil {
					exceptionCounter += 1

					// TODO ignore log
					fmt.Println("ProtoBufferUnmarshalError", err.Error())
				} else {
					msg = []byte{}

					switch upcomingMsg.Action.(type) {
					case *message.UpcomingMessage_Move_:
						if msg, err = proto.Marshal(
							&message.PublicMessage{
								Action: &message.PublicMessage_Move_{
									Move: &message.PublicMessage_Move{
										Tid: tid,
										X:   upcomingMsg.GetMove().X,
										Y:   upcomingMsg.GetMove().Y,
									},
								},
							},
						); err != nil {
							exceptionCounter += 1

							// TODO ignore log
							fmt.Println("ProtoBufferMarshalError", err.Error())
						}
					case *message.UpcomingMessage_LieDown_:
						if msg, err = proto.Marshal(
							&message.PublicMessage{
								Action: &message.PublicMessage_LieDown_{
									LieDown: &message.PublicMessage_LieDown{
										Tid: tid,
										Bed: upcomingMsg.GetLieDown().Bed,
									},
								},
							},
						); err != nil {
							exceptionCounter += 1

							// TODO ignore log
							fmt.Println("ProtoBufferMarshalError", err.Error())
						}
					default:
						exceptionCounter += 1

						// TODO ignore log
						fmt.Println("No matching UpcomingMessage")
					}

					if len(msg) > 0 {
						go plaza.Broadcast(msg)
					}
				}
			}
		}

		plaza.Leave(key)

		if msgLeave, err := proto.Marshal(
			&message.PublicMessage{
				Action: &message.PublicMessage_Leave_{
					Leave: &message.PublicMessage_Leave{
						Tid: tid,
					},
				},
			},
		); err != nil {
			fmt.Println("ProtoBufferMarshalError", err.Error())
		} else {
			go plaza.Broadcast(msgLeave)
		}
	}))

	// * Redis Bridge
	go func() {
		pubsub := rdc.Subscribe(context.Background(), plazaKey)
		defer pubsub.Close()

		bridge := channel.NewRedisBridge(rdc, pubsub, plazaKey)
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
		log.Fatalln("RedisBridgeError", err.Error())
	}()

	address := ":" + fmt.Sprint(port)
	if dev {
		address = "localhost:" + fmt.Sprint(port)
	}
	log.Fatal(app.Listen(address))
}
