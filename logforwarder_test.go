package main

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LogForwarderTestSuite struct {
	suite.Suite
}

func TestLogForwarderTestSuite(t *testing.T) {
	suite.Run(t, new(LogForwarderTestSuite))
}

func (suite *LogForwarderTestSuite) TestLogContext() {

	logContextMap := getLogContextValues("client.id")
	suite.Len(logContextMap, 2)
}

func (suite *LogForwarderTestSuite) TestMessageForwarding() {

	logForwarder := newLogForwarder(loggerForTest())
	logMessage := LogMessage{
		Message:   "log-message",
		LogLevel:  "info",
		ThingName: "aws-iot-client-id",
	}
	logForwarder.ForwardLogMessage(context.Background(), logMessage)
}

func (suite *LogForwarderTestSuite) TestBootstrap() {

	handler := bootstrap(loadConfigForTest(nil))
	suite.NotNil(handler)
}
