package handlers

import (
	"net/http"
	"encoding/json"
	"projecthub-backend/models"
)

/*
func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
    type request struct { ID string `json:"id"` }
    var req request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.Pool.Exec(context.Background(), "INSERT INTO rooms(id) VALUES($1)", req.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}*/

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
    var req struct{ ID string `json:"id"` }
    json.NewDecoder(r.Body).Decode(&req)
    models.RoomsMu.Lock()
    models.Rooms[req.ID] = models.NewRoom(req.ID)
    models.RoomsMu.Unlock()
    w.WriteHeader(http.StatusCreated)
}
