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
func (forwarder *LogForwarder) ForwardLogMessage(ctx context.Context, logMessage LogMessage) {

	logger := forwarder.loggerWithContext(ctx, logMessage.ThingName)
	defer logger.Flush()

	logger.Log(log.LogLevelByName(logMessage.LogLevel), logMessage.Message)
}

// loggerWithContext creates a log context map with given client id and some lambda context values.
func (forwarder *LogForwarder) loggerWithContext(ctx context.Context, clientId string) log.Logger {
	contextValues := make(map[string]string)
	contextValues["clientid"] = clientId
	contextValues[log.LogCtxNamespace] = "iot-display"
	return log.AppendFromLambdaContext(log.AppendContextValues(forwarder.logger, contextValues), ctx)
}
