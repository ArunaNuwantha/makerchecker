package main

import (
	"log"
	"net/http"

	"makerchecker/handlers"
	"makerchecker/router"
	"makerchecker/store"
)

func main() {
	// Initialize the in-memory store
	store := store.NewMessageStore()

	// Initialize handlers with the store dependency
	handlers := handlers.NewHandlers(store)

	// Setup the router with handlers
	router := router.SetupRouter(handlers)

	// Start the HTTP server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
