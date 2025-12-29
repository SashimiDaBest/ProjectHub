package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "projecthub-backend/db"
)

func CreateClientHandler(w http.ResponseWriter, r *http.Request) {
    type request struct { ID string `json:"id"` }
    var req request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.Pool.Exec(context.Background(), "INSERT INTO clients(id) VALUES($1)", req.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}
