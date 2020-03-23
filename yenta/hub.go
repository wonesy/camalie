package yenta

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Spoke interface {
	ID() uuid.UUID
	Send(interface{}) error
}

// Hub is the central control for a single msg channel
type Hub struct {
	ID         uuid.UUID
	spokes     []Spoke
	broadcast  chan interface{}
	register   chan Spoke
	unregister chan uuid.UUID
	stop       chan struct{}
	ended      chan struct{}
}

func NewHub() *Hub {
	id, _ := uuid.NewUUID()
	return &Hub{
		ID: id,
	}
}

func (h *Hub) Stop() error {
	h.stop <- struct{}{}

	select {
	case <-h.ended:
		return nil
	case <-time.After(5 * time.Second):
		return errors.New("timeout on stop")
	}
}

func (h *Hub) Register(sp Spoke) {
	h.register <- sp
}

func (h *Hub) Unregister(id uuid.UUID) {
	h.unregister <- id
}

func (h *Hub) Broadcast(msg interface{}) {
	h.broadcast <- msg
}

// Start ...
func (h *Hub) Start() {
	h.register = make(chan Spoke, 5)
	h.unregister = make(chan uuid.UUID, 5)
	h.stop = make(chan struct{})
	h.ended = make(chan struct{})
	h.broadcast = make(chan interface{}, 10)

	go func() {
		defer func() {
		}()

		for {
			select {
			case msg := <-h.broadcast:
				for _, sp := range h.spokes {
					if err := sp.Send(msg); err != nil {
						// TODO log this
					}
				}
			case s := <-h.register:
				alreadyRegistered := false
				for _, sp := range h.spokes {
					if s.ID() == sp.ID() {
						alreadyRegistered = true
						break
					}
				}
				if !alreadyRegistered {
					h.spokes = append(h.spokes, s)
				}
			case id := <-h.unregister:
				for i, sp := range h.spokes {
					if id == sp.ID() {
						h.spokes = append(h.spokes[:i], h.spokes[i+1:]...)
						break
					}
				}
			case <-h.stop:
				h.ended <- struct{}{}
				return
			}
		}
	}()
}
