package main

import (
	"game-of-life/api"
	"game-of-life/config"
	"game-of-life/handlers"
	"log"
	"net/http"
)

func main() {

	config.InitDB()
	defer config.DB.Close()

	go handlers.HandleMessages()
	router := api.InitRouter()

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Println("an error occured : ", err)
	}

}
