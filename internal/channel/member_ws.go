package channel

import (
	"log"
	"sync"

	"github.com/gofiber/websocket/v2"
)

type WebSocketMember struct {
	Kicked     bool
	Conn       *websocket.Conn
	writeMutex sync.Mutex
}

func (m *WebSocketMember) Receive() (msg []byte, err error) {
	if _, msg, err = m.Conn.ReadMessage(); err != nil {
		if !websocket.IsUnexpectedCloseError(err) {
			log.Println("UnexpectedCloseError:", err.Error())
		}
		return nil, err
	}
	return msg, err
}

func (m *WebSocketMember) Reply(msg []byte) {
	// ! Avoid WebSocket Concurrent Write

	m.writeMutex.Lock()
	{
		if err := m.Conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
			log.Println("WriteMessageError:", err.Error())
		}
	}
	m.writeMutex.Unlock()
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
