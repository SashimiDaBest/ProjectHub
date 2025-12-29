package models

import (
	"fmt"
	"sync"
	"log"
)

type Room struct {
    ID        string
    Join      chan *Client
    Leave     chan *Client
    Broadcast chan Message
    Clients   map[string]*Client
	mu        sync.Mutex
}

// Create
func NewRoom(id string) *Room {
    r := &Room{
        ID:        id,
        Clients:   make(map[string]*Client),
        Broadcast: make(chan Message, 256),
        Join:      make(chan *Client),
        Leave:     make(chan *Client),
    }
    go r.Run()
    return r
}

func (r *Room) Run() {
    for {
        select {
        case c := <-r.Join:
            r.mu.Lock()
            r.Clients[c.ID] = c  // store client keyed by ID
            r.mu.Unlock()

        case c := <-r.Leave:
            r.mu.Lock()
            delete(r.Clients, c.ID)  // remove by ID
            r.mu.Unlock()

        case msg := <-r.Broadcast:
            r.mu.Lock()
            for _, c := range r.Clients {  // iterate over values
                select {
                case c.Send <- msg:
                default:
                    log.Println("Dropped message for slow client:", c.ID)
                }
            }
            r.mu.Unlock()
        }
    }
}

// Read
func (r *Room) Info() string {
    clients := make([]string, 0, len(r.Clients))
    for id := range r.Clients {
        clients = append(clients, id)
    }
    return fmt.Sprintf("Room ID: %s, Clients: %v", r.ID, clients)
}

// Update (join/leave clients)
func (r *Room) AddClient(c *Client) {
    r.Clients[c.ID] = c
    c.JoinRoom(r.ID)
}
func (r *Room) RemoveClient(c *Client) {
    delete(r.Clients, c.ID)
    c.LeaveRoom(r.ID)
}

// Delete (close room)
func (r *Room) Close() {
    close(r.Join)
    close(r.Leave)
    close(r.Broadcast)
    r.Clients = nil
}

/*
// Room Goroutine 
func (r *Room) Run() {
    for {
        select {
        case c, ok := <-r.Join:
            if !ok {
                return
            }
            r.AddClient(c)
        case c, ok := <-r.Leave:
            if !ok {
                return
            }
            r.RemoveClient(c)
        case msg, ok := <-r.Broadcast:
            if !ok {
                return
            }
            for _, c := range r.Clients {
                select {
                case c.Send <- msg:
                default:
                    // slow consumer handling
                    fmt.Println("Message dropped for", c.ID)
                }
            }
        }
    }
}
*/