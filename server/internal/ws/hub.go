package ws

import (
	"fmt"
	"time"
)

type Room struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Clients     map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			r, ok := h.Rooms[cl.RoomID]
			if !ok {
				r = &Room{
					ID:          cl.RoomID,
					Name:        "",
					Description: "",
					Clients:     make(map[string]*Client),
				}
				h.Rooms[cl.RoomID] = r
			}
			r.Clients[cl.ID] = cl
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:   fmt.Sprintf("%s has left the room", cl.Username),
							CreatedAt: time.Now(),
							RoomID:    cl.RoomID,
							Username:  cl.Username,
						}
					}
					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}
		case msg := <-h.Broadcast:
			if _, ok := h.Rooms[msg.RoomID]; ok {
				for _, cl := range h.Rooms[msg.RoomID].Clients {
					cl.Message <- msg
				}
			}
		}
	}
}
