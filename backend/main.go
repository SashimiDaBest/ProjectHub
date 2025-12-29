package main

import (
    // "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"projecthub-backend/auth"
    "log"
    "projecthub-backend/db"
    "projecthub-backend/handlers"
    // "projecthub-backend/models"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from your frontend dev server
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func main() {
    // connect to db
    db.Connect()
    defer db.Pool.Close()
		
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to ProjectHub Backend"))
	}).Methods("GET")

	r.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	r.HandleFunc("/protected", auth.ProtectedHandler).Methods("GET")

    r.HandleFunc("/clients", handlers.CreateClientHandler).Methods("POST")
    r.HandleFunc("/rooms", handlers.CreateRoomHandler).Methods("POST")
    r.HandleFunc("/messages", handlers.SendMessageHandler).Methods("POST")
    r.HandleFunc("/offline", handlers.GetOfflineMessagesHandler).Methods("GET")
	r.HandleFunc("/ws", handlers.WebSocketHandler)
    r.HandleFunc("/rooms", handlers.CreateRoomHandler).Methods("POST")

	handler := enableCORS(r)

	log.Println("Starting the server on http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", handler))
}
