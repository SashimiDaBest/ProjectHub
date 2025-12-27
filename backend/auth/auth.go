package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
   // "log"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("secret-key") // keep secret and unexported

// User struct for login
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// createToken generates a JWT token for a username (private)
func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		})
	return token.SignedString(secretKey)
}

// verifyToken parses and validates the JWT
func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// LoginHandler handles user login and returns a JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
   
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request payload")
		return
	}

	// Dummy credentials check
	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := createToken(u.Username) // private function used internally
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Could not create token")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, "Invalid credentials")
}

// ProtectedHandler is an example protected endpoint
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing or invalid authorization header")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if err := verifyToken(tokenString); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to the protected area!")
}
