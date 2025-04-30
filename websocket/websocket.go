package websocket

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

type Hub struct {
	clients   map[*Client]bool
	broadcast chan []byte
	mutex     sync.Mutex
}

var GlobalHub = NewHub()

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (hub *Hub) HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	client := &Client{
		Conn: ws,
		Send: make(chan []byte),
	}

	hub.mutex.Lock()
	hub.clients[client] = true
	hub.mutex.Unlock()

	go hub.readPump(client)
	go hub.writePump(client)
}

func (hub *Hub) readPump(client *Client) {
	defer func() {
		hub.mutex.Lock()
		delete(hub.clients, client)
		hub.mutex.Unlock()
		client.Conn.Close()
	}()

	for {
		_, _, err := client.Conn.ReadMessage()
		if err != nil {
			break // ignore input; close on error
		}
	}
}

func (hub *Hub) writePump(client *Client) {
	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}

// Start runs the broadcast dispatcher
func (hub *Hub) Start() {
	for {
		msg := <-hub.broadcast
		hub.mutex.Lock()
		for client := range hub.clients {
			select {
			case client.Send <- msg:
			default:
				close(client.Send)
				delete(hub.clients, client)
			}
		}
		hub.mutex.Unlock()
	}
}
