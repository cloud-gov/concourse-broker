package logger

import (
	"code.cloudfoundry.org/lager"
	"fmt"
	"github.com/18F/concourse-broker/config"
	"os"
)

// taken from https://github.com/alphagov/paas-rds-broker/blob/3c0f2e0e6f8b8f08d95446c81cf9a43150414185/main.go
var logLevels = map[string]lager.LogLevel{
	"DEBUG": lager.DEBUG,
	"INFO":  lager.INFO,
	"ERROR": lager.ERROR,
	"FATAL": lager.FATAL,
}

func NewLogger(component string, env config.Env) (lager.Logger, error) {
	logger := lager.NewLogger(component)
	logLevel, ok := logLevels[env.LogLevel]
	if !ok {
		return nil, fmt.Errorf("Unknown log level %s. Available log levels are: DEBUG, INFO, ERROR, and FATAL", env.LogLevel)
	}
	fmt.Printf("Using log level %s\n", env.LogLevel)
	logger.RegisterSink(lager.NewWriterSink(os.Stderr, logLevel))
	return logger, nil
}
