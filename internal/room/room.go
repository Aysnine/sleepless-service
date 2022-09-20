package room

type Room struct {
	increment int
	members   map[int]Member
	bridge    Bridge
}

func New() *Room {
	room := &Room{
		increment: 0,
		members:   make(map[int]Member),
		bridge:    nil,
	}
	return room
}

func (r *Room) Join(m Member) int {
	key := r.increment + 1

	r.increment = key
	r.members[r.increment] = m

	return key
}

func (r *Room) Leave(key int) {
	delete(r.members, key)
}

func (r *Room) Broadcast(msg []byte) {
	if r.bridge != nil {
		r.bridge.Reply(msg)
	} else {
		r.Delivery(msg)
	}
}

func (r *Room) Delivery(msg []byte) {
	for _, member := range r.members {
		go member.Reply(msg)
	}
}

func (r *Room) SetBridge(bridge Bridge) {
	r.bridge = bridge
}

func (r *Room) RemoveBridge() {
	r.bridge = nil
}
