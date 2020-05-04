package chat

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
	pbchat "github.com/wonesy/camalie/proto/chat"
)

// Hub main control for chatting and message broadcasting
type Hub struct {
	ID string

	clients    map[string]*Client
	register   chan *Client
	unregister chan string
	broadcast  chan *pbchat.Message
	stop       chan struct{}
	ended      chan struct{}
}

// NewHub constuctor for Hub
func NewHub() *Hub {
	id := ksuid.New().String()
	return &Hub{
		ID: id,
	}
}

// GetID returns the hub's ID
func (h *Hub) GetID() string {
	return h.ID
}

// Register allows a client to register to the hub
func (h *Hub) Register(c *Client) {
	h.register <- c
}

// Unregister allows a client to unregister from the hub
func (h *Hub) Unregister(id string) {
	h.unregister <- id
}

// Broadcast sends a message to all other clients on the hub
func (h *Hub) Broadcast(msg *pbchat.Message) {
	h.broadcast <- msg
}

// GracefulStop stops the hub's loop
func (h *Hub) GracefulStop() {
	h.stop <- struct{}{}

	select {
	case <-h.ended:
		fmt.Println("Hub was sucessfully stopped")
	case <-time.After(10 * time.Second):
		fmt.Println("Hub was unable to gracefully stop. Force quitting")
	}
}

// Start main processing loop for chats
func (h *Hub) Start() {
	h.register = make(chan *Client)
	h.unregister = make(chan string)
	h.broadcast = make(chan *pbchat.Message, 10)
	h.stop = make(chan struct{})
	h.ended = make(chan struct{})

	for {
		select {
		case c := <-h.register:
			h.clients[c.ID()] = c
		case id := <-h.unregister:
			delete(h.clients, id)
		case msg := <-h.broadcast:
			for c, spk := range h.clients {
				if err := spk.FromHub(msg); err != nil {
					delete(h.clients, c)
				}
			}
		case <-h.stop:
			h.ended <- struct{}{}
			return
		}
	}
}
