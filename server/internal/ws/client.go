package ws

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	RoomID    string    `json:"room_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

// WriteMessage continuously sends messages to the client's WebSocket connection.
// It handles connection cleanup by closing the connection when the method exits.
// The method runs in a blocking loop until the client's message channel is closed.
//
// Parameters:
//   - None
func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		msg, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(msg)
	}
}

// ReadMessage continuously reads messages from the client's WebSocket connection
// and broadcasts them to all other clients in the same room through the hub.
// It handles connection cleanup by unregistering the client from the hub when
// the connection is closed or encounters an error. The method runs in a blocking
// loop until the WebSocket connection is terminated.
//
// Parameters:
// hub: The Hub instance used to broadcast messages and handle client unregistration
//
// The method will:
//   - Automatically unregister the client and close the connection on function exit
//   - Log unexpected WebSocket close errors for debugging purposes
//   - Convert raw message bytes to a structured Message with metadata
//   - Send each received message to the hub's broadcast channel
func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message := &Message{
			Content:   string(msg),
			CreatedAt: time.Now(),
			RoomID:    c.RoomID,
			Username:  c.Username,
		}

		hub.Broadcast <- message
	}
}
