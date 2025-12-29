package models

import (
    "golang.org/x/time/rate"
    "fmt"
    "sync"
    "github.com/gorilla/websocket"
    "log"
    "time"
    "context"
    "projecthub-backend/db"
)
var (
    Rooms   = make(map[string]*Room) // global room map
    RoomsMu sync.Mutex               // mutex to protect concurrent access
)

type Client struct {
    ID       string
    Send     chan Message // outbound
    Conn     *websocket.Conn
    Rooms    map[string]bool
    Rate     *rate.Limiter
    Password string // for now
}

// Create
func NewClient(id string) *Client {
    return &Client{
        ID:    id,
        Send:  make(chan Message, 64),
        Rooms: make(map[string]bool),
        Rate:  rate.NewLimiter(5, 10), // 5 msgs/sec, burst 10
    }
}

// Read
func (c *Client) Info() string {
    rooms := make([]string, 0, len(c.Rooms))
    for r := range c.Rooms {
        rooms = append(rooms, r)
    }
    return fmt.Sprintf("Client ID: %s, Rooms: %v", c.ID, rooms)
}

// Update (join/leave rooms)
func (c *Client) JoinRoom(roomID string) {
    c.Rooms[roomID] = true
}
func (c *Client) LeaveRoom(roomID string) {
    delete(c.Rooms, roomID)
}

// Delete (disconnect client) -- still unser if i want temp client or nah
func (c *Client) Disconnect() {
    close(c.Send)
    c.Rooms = nil
}

func (c *Client) WritePump() {
    for msg := range c.Send {
        c.Conn.WriteJSON(msg)
    }
}

func (c *Client) ReadPump() {
    defer c.Conn.Close()
    for {
        var msg Message
        if err := c.Conn.ReadJSON(&msg); err != nil {
            log.Println("read error:", err)
            return
        }

        // Rate limit
        if !c.Rate.Allow() {
            continue
        }

        msg.Timestamp = time.Now()
        // Save message to DB
        err := db.Pool.QueryRow(context.Background(),
            "INSERT INTO messages(from_client, room_id, body, timestamp) VALUES($1,$2,$3,$4) RETURNING id",
            msg.From, msg.RoomID, msg.Body, msg.Timestamp).Scan(&msg.ID)
        if err != nil {
            log.Println("DB insert error:", err)
        }

        // Broadcast to room
        RoomsMu.Lock()
        Room, ok := Rooms[msg.RoomID]
        RoomsMu.Unlock()
        if ok {
            Room.Broadcast <- msg
        }
    }
}