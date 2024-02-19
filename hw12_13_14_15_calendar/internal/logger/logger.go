package logger

import (
	log "github.com/sirupsen/logrus"
)

type ILog interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type Logger struct {
	ILog
	stdOutLog *log.Logger
}

func (l *Logger) Debug(args ...interface{}) {
	if l.stdOutLog != nil {
		l.stdOutLog.Debug(args)
	}
}

func (l *Logger) Info(args ...interface{}) {
	if l.stdOutLog != nil {
		l.stdOutLog.Info(args)
	}
}

func (l *Logger) Error(args ...interface{}) {
	if l.stdOutLog != nil {
		l.stdOutLog.Error(args)
	}
}

func (l *Logger) Fatal(args ...interface{}) {
	if l.stdOutLog != nil {
		l.stdOutLog.Fatal(args)
	}
}

func (l *Logger) Panic(args ...interface{}) {
	if l.stdOutLog != nil {
		l.stdOutLog.Panic(args)
	}
}
