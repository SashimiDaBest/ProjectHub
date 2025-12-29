package handlers

import (
	"net/http"
	"context"
	"projecthub-backend/models"
	"projecthub-backend/db"
	"encoding/json"
)

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
    type request struct {
        From   string `json:"from"`
        RoomID string `json:"room_id"`
        Body   string `json:"body"`
    }
    var req request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var messageID int
    err := db.Pool.QueryRow(context.Background(),
        "INSERT INTO messages(from_client, room_id, body) VALUES($1,$2,$3) RETURNING id",
        req.From, req.RoomID, req.Body,
    ).Scan(&messageID)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Insert into offline_messages for all clients in room except sender
    _, err = db.Pool.Exec(context.Background(), `
        INSERT INTO offline_messages(client_id, message_id)
        SELECT client_id, $1 FROM client_rooms WHERE room_id=$2 AND client_id<>$3
    `, messageID, req.RoomID, req.From)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func GetOfflineMessagesHandler(w http.ResponseWriter, r *http.Request) {
    clientID := r.URL.Query().Get("client_id")
    if clientID == "" {
        http.Error(w, "missing client_id", http.StatusBadRequest)
        return
    }

    rows, err := db.Pool.Query(context.Background(), `
        SELECT m.id, m.from_client, m.room_id, m.body, m.timestamp
        FROM messages m
        JOIN offline_messages o ON m.id = o.message_id
        WHERE o.client_id=$1 AND o.delivered=false
    `, clientID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    messages := []models.Message{}
    for rows.Next() {
        var m models.Message
        if err := rows.Scan(&m.ID, &m.From, &m.RoomID, &m.Body, &m.Timestamp); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        messages = append(messages, m)
    }

    // mark as delivered
    _, _ = db.Pool.Exec(context.Background(), "UPDATE offline_messages SET delivered=true WHERE client_id=$1", clientID)

    json.NewEncoder(w).Encode(messages)
}
