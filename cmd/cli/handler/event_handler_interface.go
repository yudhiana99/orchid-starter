package handler

import (
	"context"

	"github.com/mataharibiz/sange/v2"
)

// EventHandlerInterface defines the contract for event handlers
type EventHandlerInterface interface {
	Handle(ctx context.Context, event sange.EventData) error
	GetEventTypes() []string
}

// EventHandlerConfig holds configuration for event handlers
type EventHandlerConfig struct {
	Name         string
	Alias        string
	Usage        string
	Description  string
	QueueName    string
	ExchangeName string
	LoggerPrefix string
}
