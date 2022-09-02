package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"
)

func main() {

	handler := bootstrap(nil)
	lambda.Start(handler.ForwardLogMessage)
}

// bootstrap loads config and creates a log handler.
func bootstrap(conf config.Config) Handler {

	if conf == nil {
		conf = loadConfig()
	}
	secretsManager := newSecretsManager()
	return newLogForwarder(newLogger(conf, secretsManager))
}

// loadConfig from config file.
func loadConfig() config.Config {

	configSource, err := config.NewS3ConfigSourceFromEnv()
	if err != nil {
		panic(err)
	}

	conf, err := configSource.Load()
	if err != nil {
		panic(err)
	}
	return conf
}

// newSecretsManager retruns a new secrets manager from passed config.
func newSecretsManager() secrets.SecretsManager {
	return secrets.NewSecretsManager()
}

// newLogger creates a new logger from  passed config.
func newLogger(conf config.Config, secretsMenager secrets.SecretsManager) log.Logger {
	logger := log.NewLoggerFromConfig(conf, secretsMenager)
	return log.WithNameSpace(logger, "iot-display")
}
