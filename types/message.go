package types

import "encoding/json"

// Message represents the structure of messages exchanged with the extension
type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
