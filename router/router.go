package router

import (
	"makerchecker/handlers"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the router with all the routes and their handlers.
func SetupRouter(h *handlers.Handlers) *mux.Router {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/messages", h.CreateMessageHandler).Methods("POST")
	router.HandleFunc("/messages/{id}/approve", h.ApproveMessageHandler).Methods("POST")
	router.HandleFunc("/messages/{id}/reject", h.RejectMessageHandler).Methods("POST")
	router.HandleFunc("/messages", h.ListMessagesHandler).Methods("GET")

	return router
}
