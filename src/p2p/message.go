package p2p

import (
	"encoding/json"
	"fmt"
)

// Message represents a message sent between nodes.
type Message struct {
	Type    string      `json:"type"`    // Type of message (e.g., "transaction", "block")
	Payload interface{} `json:"payload"` // Payload of the message
}

// Serialize serializes the message to JSON.
func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

// Deserialize deserializes the message from JSON.
func Deserialize(data []byte) (*Message, error) {
	var msg Message
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("failed to deserialize message: %v", err)
	}
	return &msg, nil
}
