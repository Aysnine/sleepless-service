package room

type Bridge interface {
	Receive() (msg []byte, err error)
	Reply(msg []byte)
}
