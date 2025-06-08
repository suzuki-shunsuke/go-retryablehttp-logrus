// Package rlog provides a hashicorp/go-retryablehttp's LeveledLogger for sirupsen/logrus.
package rlog

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Entry is an interface to abstract logrus.Entry.
type Entry interface {
	WithFields(fields logrus.Fields) *logrus.Entry
}

// Logger is a logger implementing hashicorp/go-retryablehttp's LeveledLogger interface.
type Logger struct {
	entry      Entry
	debugLevel logrus.Level
	errorLevel logrus.Level
	infoLevel  logrus.Level
	warnLevel  logrus.Level
}

// New creates a new Logger.
func New(entry Entry) *Logger {
	return &Logger{
		entry:      entry,
		debugLevel: logrus.DebugLevel,
		errorLevel: logrus.ErrorLevel,
		infoLevel:  logrus.InfoLevel,
		warnLevel:  logrus.WarnLevel,
	}
}

// ChangeDebugLevel changes the log level of logger.Debug().
func (l *Logger) ChangeDebugLevel(to logrus.Level) {
	l.debugLevel = to
}

// ChangeErrorLevel changes the log level of logger.Error().
func (l *Logger) ChangeErrorLevel(to logrus.Level) {
	l.errorLevel = to
}

// ChangeInfoLevel changes the log level of logger.Info().
func (l *Logger) ChangeInfoLevel(to logrus.Level) {
	l.infoLevel = to
}

// ChangeWarnLevel changes the log level of logger.Warn().
func (l *Logger) ChangeWarnLevel(to logrus.Level) {
	l.warnLevel = to
}

// Debug outputs a debug log.
func (l *Logger) Debug(msg string, keysAndValues ...any) {
	l.log(l.debugLevel, msg, keysAndValues...)
}

// Error outputs an error log.
func (l *Logger) Error(msg string, keysAndValues ...any) {
	l.log(l.errorLevel, msg, keysAndValues...)
}

// Info outputs an info log.
func (l *Logger) Info(msg string, keysAndValues ...any) {
	l.log(l.infoLevel, msg, keysAndValues...)
}

// Warn outputs a warn log.
func (l *Logger) Warn(msg string, keysAndValues ...any) {
	l.log(l.warnLevel, msg, keysAndValues...)
}

func (l *Logger) log(level logrus.Level, msg string, keysAndValues ...any) {
	l.entry.WithFields(createFields(keysAndValues)).Log(level, msg)
}

func createFields(keysAndValues []any) logrus.Fields {
	s := len(keysAndValues)
	if s%2 != 0 {
		s--
	}
	cnt := s / 2 //nolint:mnd
	fields := make(logrus.Fields, cnt)
	for i := range cnt {
		fields[fmt.Sprint(keysAndValues[i*2])] = keysAndValues[i*2+1]
	}
	return fields
}
