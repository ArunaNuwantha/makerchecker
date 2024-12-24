package models

// MessageStatus represents the current status of a message request.
type MessageStatus string

const (
	StatusPending  MessageStatus = "Pending"
	StatusApproved MessageStatus = "Approved"
	StatusRejected MessageStatus = "Rejected"
	StatusSent     MessageStatus = "Sent"
)

// Message represents a message send request.
type Message struct {
	ID        string        `json:"id"`
	Recipient string        `json:"recipient"`
	Content   string        `json:"content"`
	Status    MessageStatus `json:"status"`
}
