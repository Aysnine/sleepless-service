package room

type Room struct {
	increment int
	members   map[int]Member
}

func New() *Room {
	room := &Room{
		increment: 0,
		members:   make(map[int]Member),
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
	for _, member := range r.members {
		go member.Reply(msg)
	}
}

func (r *Room) Delivery(msg []byte, selfKey int) {
	for key, member := range r.members {
		if key != selfKey {
			go member.Reply(msg)
		}
	}
}
