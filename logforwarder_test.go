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

	logContextMap := logContext("topic", "client.id")
	suite.Len(logContextMap, 2)
}

func (suite *LogForwarderTestSuite) TestMessageForwarding() {

	logForwarder := newLogForwarder(loggerForTest())

	logForwarder.ForwardLogMessage(context.Background(), "log-message", "info", "mqtt-tipic", "aws-iot-client-id")
}

func (suite *LogForwarderTestSuite) TestBootstrap() {

	handler := bootstrap(loadConfigForTest(nil))
	suite.NotNil(handler)
}
