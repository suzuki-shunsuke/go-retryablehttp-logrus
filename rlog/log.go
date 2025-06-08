package rlog

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Entry interface {
	WithFields(fields logrus.Fields) *logrus.Entry
}

type Logger struct {
	entry      Entry
	debugLevel logrus.Level
	errorLevel logrus.Level
	infoLevel  logrus.Level
	warnLevel  logrus.Level
}

func New(entry Entry) *Logger {
	return &Logger{
		entry:      entry,
		debugLevel: logrus.DebugLevel,
		errorLevel: logrus.ErrorLevel,
		infoLevel:  logrus.InfoLevel,
		warnLevel:  logrus.WarnLevel,
	}
}

func (l *Logger) ChangeDebugLevel(to logrus.Level) {
	l.debugLevel = to
}

func (l *Logger) ChangeErrorLevel(to logrus.Level) {
	l.errorLevel = to
}

func (l *Logger) ChangeInfoLevel(to logrus.Level) {
	l.infoLevel = to
}

func (l *Logger) ChangeWarnLevel(to logrus.Level) {
	l.warnLevel = to
}

func (l *Logger) Debug(msg string, keysAndValues ...any) {
	l.log(l.debugLevel, msg, keysAndValues...)
}

func (l *Logger) Error(msg string, keysAndValues ...any) {
	l.log(l.errorLevel, msg, keysAndValues...)
}

func (l *Logger) Info(msg string, keysAndValues ...any) {
	l.log(l.infoLevel, msg, keysAndValues...)
}

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
