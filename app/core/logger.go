package core

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Log struct {
	logger *logrus.Logger
}

func (log *Log) Infof(format string, args ...interface{}) {
	log.logger.Infof(format, args)
}

func (log *Log) Warnf(format string, args ...interface{}) {
	log.logger.Warnf(format, args)
}

func (log *Log) Errorf(format string, args ...interface{}) {
	log.logger.Errorf(format, args)
}

func (log *Log) Fatalf(format string, args ...interface{}) {
	log.logger.Fatalf(format, args)
}

func NewLogrusLogger() *logrus.Logger {
	var logger = logrus.New()
	logger.Formatter = new(logrus.TextFormatter)
	logger.Level = logrus.TraceLevel
	logger.Out = os.Stdout
	return logger
}

func NewLogger(logger *logrus.Logger) *Log {
	return &Log{
		logger: logger,
	}
}
