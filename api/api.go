package api

import (
	"game-of-life/handlers"
	"game-of-life/middleware"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

	r := mux.NewRouter()

	// Websocket routes
	r.HandleFunc("/ws", handlers.HandleWebsocket)

	// Public routes
	r.HandleFunc("/register", handlers.HandleRegister).Methods("POST")
	r.HandleFunc("/login", handlers.HandleLogin).Methods("POST")

	// Protected routes (using AuthMiddleware)
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.HandleFunc("/rooms", handlers.GetRooms).Methods("GET")

	return r
}
