package store

import (
	"makerchecker/models"
	"sync"
)

// MessageStore provides thread-safe access to a map of Messages.
type MessageStore struct {
	sync.RWMutex
	messages map[string]*models.Message
}

// NewMessageStore initializes a new MessageStore.
func NewMessageStore() *MessageStore {
	return &MessageStore{
		messages: make(map[string]*models.Message),
	}
}

// AddMessage adds a new message to the store.
func (s *MessageStore) AddMessage(msg *models.Message) {
	s.Lock()
	defer s.Unlock()
	s.messages[msg.ID] = msg
}

// GetMessage retrieves a message by ID.
func (s *MessageStore) GetMessage(id string) (*models.Message, bool) {
	s.RLock()
	defer s.RUnlock()
	msg, exists := s.messages[id]
	return msg, exists
}

// UpdateMessage updates an existing message.
func (s *MessageStore) UpdateMessage(msg *models.Message) {
	s.Lock()
	defer s.Unlock()
	s.messages[msg.ID] = msg
}

// ListMessages returns all messages.
func (s *MessageStore) ListMessages() []*models.Message {
	s.RLock()
	defer s.RUnlock()
	messages := make([]*models.Message, 0, len(s.messages))
	for _, msg := range s.messages {
		messages = append(messages, msg)
	}
	return messages
}
