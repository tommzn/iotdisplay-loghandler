package main

import (
	log "github.com/tommzn/go-log"
)

// LogMessage is a single log message including it's log level
// and client id/thing name from AWS IOT.
type LogMessage struct {
	Message   string `json:"message"`
	LogLevel  string `json:"loglevel"`
	ThingName string `json:"thing_name"`
}

// LogForwarder will send messages to a log target.
type LogForwarder struct {
	logger log.Logger
}
