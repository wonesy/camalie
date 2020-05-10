package chat

import (
	"github.com/segmentio/ksuid"
	pbchat "github.com/wonesy/camalie/proto/chat"
)

// ClientWebSocket implements a websocket
type ClientWebSocket interface {
	ReadJSON(v interface{}) error
	WriteJSON(v interface{}) error
}

// Client represents a chat client
type Client struct {
	id string
	ws ClientWebSocket
	h  *Hub
}

// NewClient constructor for Client
func NewClient(ws ClientWebSocket) *Client {
	return &Client{
		id: ksuid.New().String(),
		ws: ws,
	}
}

// JoinHub allows the client to join a new chat hub
func (c *Client) JoinHub(h *Hub) {
	c.h = h
}

// LeaveHub has a client leave a specific hub
func (c *Client) LeaveHub(h *Hub) {
	h.Unregister(c.ID())
}

// ID returns the client's ID
func (c *Client) ID() string {
	return c.id
}

// FromHub reads messages from the hub and sends them along the ws
func (c *Client) FromHub(msg *pbchat.Message) error {
	return c.ws.WriteJSON(msg)
}

// Start launches the client loop that send messages client -> hub
func (c *Client) Start() {
	go func() {
		for {
			var msg *pbchat.Message
			if err := c.ws.ReadJSON(msg); err != nil {
				// TODO log
				continue
			}

			c.h.Broadcast(msg)
		}
	}()
}
