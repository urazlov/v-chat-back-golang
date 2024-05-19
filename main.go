package main

import (
	"log"
	"v-chat-back-go/database"
	"v-chat-back-go/router"
)

func main() {
	database.Connect();
	
	defer database.Close()

	r := router.SetupRouter()
	if err := r.Run(":3100"); err != nil {
		log.Fatalf("Failed to run Server: %v\n", err)
	}
}
