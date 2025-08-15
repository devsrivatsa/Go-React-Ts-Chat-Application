package ws

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type HubHandler struct {
	hub *Hub
}

func NewHubHandler(hub *Hub) *HubHandler {
	return &HubHandler{
		hub: hub,
	}
}

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *HubHandler) CreateRoom(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room := &Room{
		ID:          req.ID,
		Name:        req.Name,
		Description: "",
		Clients:     make(map[string]*Client),
	}

	h.hub.Rooms[req.ID] = room

	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		log.Println("origin", origin)
		// return origin == "http://localhost:3000" the frontend url.
		return true

	},
}

func (h *HubHandler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	roomID := c.Param("roomId")
	userId := c.Query("clientId")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       userId,
		RoomID:   roomID,
		Username: username,
	}

	joinNotification := &Message{
		RoomID:    roomID,
		Username:  username,
		Content:   fmt.Sprintf("a new user %s has joined the room", username),
		CreatedAt: time.Now(),
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- joinNotification

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}

type RoomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ClientResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *HubHandler) GetRooms(c *gin.Context) {
	rooms := make([]RoomResponse, 0)
	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomResponse{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

func (h *HubHandler) GetClientsInRoom(c *gin.Context) {
	roomId := c.Param("roomId")
	clients := make([]ClientResponse, 0)
	if room, ok := h.hub.Rooms[roomId]; !ok {
		clients = append(clients, ClientResponse{
			ID:       room.ID,
			Username: room.Name,
		})
		c.JSON(http.StatusOK, clients)
		return
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientResponse{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}
