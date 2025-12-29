package handlers

import (
	"log"
	"net/http"
	"projecthub-backend/models"
	"golang.org/x/time/rate"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("upgrade error:", err)
        return
    }

    client := &models.Client{
        ID:    r.URL.Query().Get("id"),
        Conn:  conn,
        Send:  make(chan models.Message, 64),
        Rate:  rate.NewLimiter(5, 10),
        Rooms: make(map[string]bool),
    }

    go client.WritePump()
    client.ReadPump()
}
