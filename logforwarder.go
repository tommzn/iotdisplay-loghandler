package main

import (
	"context"

	log "github.com/tommzn/go-log"
)

// newLogForwarder creates a log forwarder with given logger.
func newLogForwarder(logger log.Logger) *LogForwarder {
	return &LogForwarder{logger: logger}
}

// ForwardLogMessage will send passed message to ued log target.
func (forwarder *LogForwarder) ForwardLogMessage(ctx context.Context, message, logLevelAsString, clientId string) {

	forwarder.logger.WithContext(ctx)
	logger := log.AppendContextValues(forwarder.logger, logContext(clientId))
	defer logger.Flush()

	logger.Log(log.LogLevelByName(logLevelAsString), message)
}

// logContext creates a log context map with given topic and cloent id.
func logContext(clientId string) map[string]string {
	values := make(map[string]string)
	values["clientid"] = clientId
	return values
}
