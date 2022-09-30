package channel

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

type WebSocketMember struct {
	Kicked bool
	Conn   *websocket.Conn
}

func (m *WebSocketMember) Receive() (msg []byte, err error) {
	if _, msg, err = m.Conn.ReadMessage(); err != nil {
		if !websocket.IsUnexpectedCloseError(err) {
			log.Println("UnexpectedCloseError:", err)
		}
		return nil, err
	}
	return msg, err
}

func (m *WebSocketMember) Reply(msg []byte) {
	if err := m.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Println("WriteMessageError:", err)
	}
}

func (m *WebSocketMember) Kick() {
	m.Kicked = true
}

func (m *WebSocketMember) IsKicked() bool {
	return m.Kicked
}

func NewWebSocketMember(conn *websocket.Conn) *WebSocketMember {
	member := &WebSocketMember{
		Kicked: false,
		Conn:   conn,
	}
	return member
}
