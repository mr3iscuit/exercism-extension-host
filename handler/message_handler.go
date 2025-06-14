package handler

import (
	"encoding/json"
	"log"

	"github.com/mr3iscuit/exercism-extension-host/types"
)

// HandleMessage processes incoming messages and returns appropriate responses
func HandleMessage(msg []byte) ([]byte, error) {
	var message types.Message
	if err := json.Unmarshal(msg, &message); err != nil {
		return nil, err
	}

	// Handle different message types
	switch message.Type {
	case "ping":
		return json.Marshal(types.Message{
			Type:    "pong",
			Payload: message.Payload,
		})
	default:
		log.Printf("Received unknown message type: %s", message.Type)
		return json.Marshal(types.Message{
			Type:    "error",
			Payload: []byte(`{"error": "unknown message type"}`),
		})
	}
}
