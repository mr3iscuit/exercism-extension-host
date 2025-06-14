package router

import (
	native "github.com/rickypc/native-messaging-host"
)

// HandlerFunc defines a function that processes a message payload and returns a response
// The response will be automatically marshaled to JSON
type HandlerFunc func(data *native.H) (*native.H, error)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) On(messageType string, handler HandlerFunc) {
	r.handlers[messageType] = handler
}

func (r *Router) HandleMessage(msg *native.H) (*native.H, error) {
	messageType, ok := (*msg)["type"].(string)
	if !ok {
		return &native.H{
			"error": "invalid message type",
		}, nil
	}

	handler, exists := r.handlers[messageType]
	if !exists {
		return &native.H{
			"error": "unknown message type: " + messageType,
		}, nil
	}

	return handler(msg)
}
