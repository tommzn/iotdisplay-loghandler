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

	logContextMap := logContext("client.id")
	suite.Len(logContextMap, 1)
}

func (suite *LogForwarderTestSuite) TestMessageForwarding() {

	logForwarder := newLogForwarder(loggerForTest())

	logForwarder.ForwardLogMessage(context.Background(), "log-message", "info", "aws-iot-client-id")
}

func (suite *LogForwarderTestSuite) TestBootstrap() {

	handler := bootstrap(loadConfigForTest(nil))
	suite.NotNil(handler)
}
