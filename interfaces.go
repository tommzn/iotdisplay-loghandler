package main

import "context"

// Handler will forward messages to log target.
type Handler interface {

	// ForwardLogMessage sends passed messages to used log target.
	ForwardLogMessage(context.Context, LogMessage)
}
