package logger

import (
	log "github.com/sirupsen/logrus"
)

var _logger *Logger

func InitLogger(logLevel string) error {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	stdOutLog := new(log.Logger)
	stdOutLog = log.New()
	stdOutLog.Level = level

	_logger = &Logger{
		stdOutLog: stdOutLog,
	}

	return nil
}

func UseLogger() *Logger {
	return _logger
}
