package rlog

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Entry interface {
	WithFields(fields logrus.Fields) *logrus.Entry
}

type Logger struct {
	entry Entry
}

func New(entry Entry) *Logger {
	return &Logger{
		entry: entry,
	}
}

func (l *Logger) Debug(msg string, keysAndValues ...any) {
	l.entry.WithFields(createFields(keysAndValues)).Debug(msg)
}

func (l *Logger) Error(msg string, keysAndValues ...any) {
	l.entry.WithFields(createFields(keysAndValues)).Error(msg)
}

func (l *Logger) Info(msg string, keysAndValues ...any) {
	l.entry.WithFields(createFields(keysAndValues)).Info(msg)
}

func (l *Logger) Warn(msg string, keysAndValues ...any) {
	l.entry.WithFields(createFields(keysAndValues)).Warn(msg)
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
