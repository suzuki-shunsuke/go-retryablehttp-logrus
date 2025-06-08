package rlog

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Level represents the log level
type Level int

const (
	// DebugLevel logs debug messages
	DebugLevel Level = iota
	// InfoLevel logs informational messages
	InfoLevel
	// WarnLevel logs warning messages
	WarnLevel
	// ErrorLevel logs error messages
	ErrorLevel
)

// String returns the string representation of the log level
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	default:
		return "unknown"
	}
}

// ToLogrusLevel converts the Level to logrus.Level
func (l Level) ToLogrusLevel() logrus.Level {
	switch l {
	case DebugLevel:
		return logrus.DebugLevel
	case InfoLevel:
		return logrus.InfoLevel
	case WarnLevel:
		return logrus.WarnLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

type Entry interface {
	WithFields(fields logrus.Fields) *logrus.Entry
}

type Logger struct {
	entry  Entry
	levels map[Level]logrus.Level
}

func New(entry Entry) *Logger {
	return &Logger{
		entry: entry,
		levels: map[Level]logrus.Level{
			DebugLevel: logrus.DebugLevel,
			InfoLevel:  logrus.InfoLevel,
			WarnLevel:  logrus.WarnLevel,
			ErrorLevel: logrus.ErrorLevel,
		},
	}
}

func (l *Logger) ChangeLevel(from Level, to logrus.Level) {
	l.levels[from] = to
}

func (l *Logger) Debug(msg string, keysAndValues ...any) {
	l.log(DebugLevel, msg, keysAndValues...)
}

func (l *Logger) Error(msg string, keysAndValues ...any) {
	l.log(ErrorLevel, msg, keysAndValues...)
}

func (l *Logger) Info(msg string, keysAndValues ...any) {
	l.log(InfoLevel, msg, keysAndValues...)
}

func (l *Logger) Warn(msg string, keysAndValues ...any) {
	l.log(WarnLevel, msg, keysAndValues...)
}

func (l *Logger) log(level Level, msg string, keysAndValues ...any) {
	lvl, ok := l.levels[level]
	if !ok {
		lvl = level.ToLogrusLevel()
	}
	l.entry.WithFields(createFields(keysAndValues)).Log(lvl, msg)
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
