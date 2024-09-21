package go_chat

//comment
import "github.com/gorilla/websocket"

// Client
type client struct {
	socket *websocket.Conn // Client web socket
	send   chan []byte     // Send channel
	room   *room           // Client room is in
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err != nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
