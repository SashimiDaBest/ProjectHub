package main

import (
    "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"projecthub-backend/auth"
    "log"
    // "projecthub-backend/db"
    // "projecthub-backend/handlers"
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
	router := mux.NewRouter()

	// Add routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to ProjectHub Backend"))
	}).Methods("GET")

	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", auth.ProtectedHandler).Methods("GET")

	// Wrap router with CORS middleware
	handler := enableCORS(router)

	fmt.Println("Starting the server on http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", handler))

	/*
    // Connect to database
    db.Connect()
    defer db.Pool.Close()

    // Test: create a user
    user := models.User{Name: "Alice", Email: "alice@example.com"}
    if err := handlers.CreateUser(user); err != nil {
        log.Fatalf("Failed to create user: %v", err)
    }

    // Test: list users
    users, err := handlers.GetUsers()
    if err != nil {
        log.Fatalf("Failed to get users: %v", err)
    }

    fmt.Println("Users in database:")
    for _, u := range users {
        fmt.Printf("%d: %s (%s)\n", u.ID, u.Name, u.Email)
    }
	*/
}
