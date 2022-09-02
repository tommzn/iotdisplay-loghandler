package main

import (
	log "github.com/tommzn/go-log"
)

// LogForwarder will send messages to a log target.
type LogForwarder struct {
	logger log.Logger
}
