package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"makerchecker/models"
	"makerchecker/store"
)

// Handlers struct holds dependencies for HTTP handlers.
type Handlers struct {
	store *store.MessageStore
}

// NewHandlers initializes Handlers with necessary dependencies.
func NewHandlers(store *store.MessageStore) *Handlers {
	return &Handlers{
		store: store,
	}
}

// CreateMessageHandler handles the creation of a new message send request.
func (h *Handlers) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Recipient string `json:"recipient"`
		Content   string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Recipient == "" || req.Content == "" {
		http.Error(w, "Recipient and Content are required", http.StatusBadRequest)
		return
	}

	msg := &models.Message{
		ID:        uuid.New().String(),
		Recipient: req.Recipient,
		Content:   req.Content,
		Status:    models.StatusPending,
	}

	h.store.AddMessage(msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

// ApproveMessageHandler handles the approval of a message send request.
func (h *Handlers) ApproveMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	msg, exists := h.store.GetMessage(id)
	if !exists {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	if msg.Status != models.StatusPending {
		http.Error(w, "Message is not in a pending state", http.StatusBadRequest)
		return
	}

	// Update status to Approved
	msg.Status = models.StatusApproved
	h.store.UpdateMessage(msg)

	// Simulate sending the message
	// In a real-world scenario, integrate with an email/SMS service here.
	log.Printf("Sending message to %s: %s", msg.Recipient, msg.Content)
	msg.Status = models.StatusSent
	h.store.UpdateMessage(msg)

	response := struct {
		ID      string               `json:"id"`
		Status  models.MessageStatus `json:"status"`
		Message string               `json:"message"`
	}{
		ID:      msg.ID,
		Status:  msg.Status,
		Message: "Message sent successfully.",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RejectMessageHandler handles the rejection of a message send request.
func (h *Handlers) RejectMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	msg, exists := h.store.GetMessage(id)
	if !exists {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	if msg.Status != models.StatusPending {
		http.Error(w, "Message is not in a pending state", http.StatusBadRequest)
		return
	}

	// Update status to Rejected
	msg.Status = models.StatusRejected
	h.store.UpdateMessage(msg)

	response := struct {
		ID      string               `json:"id"`
		Status  models.MessageStatus `json:"status"`
		Message string               `json:"message"`
	}{
		ID:      msg.ID,
		Status:  msg.Status,
		Message: "Message send request has been rejected.",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ListMessagesHandler handles listing all message send requests.
func (h *Handlers) ListMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages := h.store.ListMessages()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
