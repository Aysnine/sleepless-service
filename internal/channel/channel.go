package channel

type Channel struct {
	increment int
	members   map[int]Member
	bridge    Bridge
}

func New() *Channel {
	channel := &Channel{
		increment: 0,
		members:   make(map[int]Member),
		bridge:    nil,
	}
	return channel
}

func (r *Channel) Join(m Member) int {
	key := r.increment + 1

	r.increment = key
	r.members[r.increment] = m

	return key
}

func (r *Channel) Leave(key int) {
	delete(r.members, key)
}

func (r *Channel) Broadcast(msg []byte) {
	if r.bridge != nil {
		r.bridge.Reply(msg)
	} else {
		r.Delivery(msg)
	}
}

func (r *Channel) Delivery(msg []byte) {
	for _, member := range r.members {
		go member.Reply(msg)
	}
}

func (r *Channel) SetBridge(bridge Bridge) {
	r.bridge = bridge
}

func (r *Channel) RemoveBridge() {
	r.bridge = nil
}
