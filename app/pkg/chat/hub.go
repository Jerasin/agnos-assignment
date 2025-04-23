package chat

import "sync"

type Hub struct {
	rooms map[string]*Room
	mu    sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string]*Room),
	}
}

func (h *Hub) GetRoom(name string) *Room {
	h.mu.Lock()
	defer h.mu.Unlock()

	room, ok := h.rooms[name]
	if !ok {
		room = NewRoom(name)
		h.rooms[name] = room
		go room.Run()
	}
	return room
}
