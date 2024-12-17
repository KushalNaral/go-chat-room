package handlers

import (
	"encoding/json"
	"fmt"
	"game-of-life/config"
	"game-of-life/models"
	"game-of-life/utils"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// Handler for user registration
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(creds.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Insert user into database
	_, err = config.DB.Exec("INSERT INTO users (id, username, password_hash, email) VALUES (UUID(), ?, ?, ?)",
		creds.Username, string(hashedPassword), creds.Email)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

// Handler for user login
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Fetch user's hashed password
	var storedPassword string
	err := config.DB.QueryRow("SELECT password_hash FROM users WHERE username = ?", creds.Username).Scan(&storedPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	// Compare passwords
	if err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	token, err := utils.GenerateToken(creds.Username, time.Hour)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Send the token to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func GetRooms(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We have these rooms"))
}
