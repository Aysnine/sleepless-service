package channel

type Member interface {
	Receive() (msg []byte, err error)
	Reply(msg []byte)
	Kick()
	IsKicked() bool
}
