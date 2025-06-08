package rlog_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-retryablehttp-logrus/rlog"
)

func newLogger(buf io.Writer) *rlog.Logger {
	l := logrus.New()
	l.Out = buf
	l.Level = logrus.DebugLevel
	l.Formatter = &logrus.TextFormatter{
		DisableTimestamp: true,
	}
	logE := logrus.NewEntry(l)
	logger := rlog.New(logE)
	return logger
}

func TestLogger_Debug(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := newLogger(buf)
	logger.Debug("retrying request", "request", "GET", "status", 500)
	s := buf.String()
	want := `level=debug msg="retrying request" request=GET status=500` + "\n"
	if s != want {
		t.Fatalf("wanted %s, got %s", want, s)
	}
}

func TestLogger_Error(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := newLogger(buf)
	logger.Error("retrying request", "request", "GET", "status", 500)
	s := buf.String()
	want := `level=error msg="retrying request" request=GET status=500` + "\n"
	if s != want {
		t.Fatalf("wanted %s, got %s", want, s)
	}
}

func TestLogger_Info(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := newLogger(buf)
	logger.Info("retrying request", "request", "GET", "status", 500)
	s := buf.String()
	want := `level=info msg="retrying request" request=GET status=500` + "\n"
	if s != want {
		t.Fatalf("wanted %s, got %s", want, s)
	}
}

func TestLogger_Warn(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := newLogger(buf)
	logger.Warn("retrying request", "request", "GET", "status", 500)
	s := buf.String()
	want := `level=warning msg="retrying request" request=GET status=500` + "\n"
	if s != want {
		t.Fatalf("wanted %s, got %s", want, s)
	}
}

func TestLogger_ChangeDebugLevel(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := newLogger(buf)
	logger.ChangeDebugLevel(logrus.InfoLevel)
	logger.Debug("retrying request", "request", "GET", "status", 500)
	s := buf.String()
	want := `level=info msg="retrying request" request=GET status=500` + "\n"
	if s != want {
		t.Fatalf("wanted %s, got %s", want, s)
	}
}

func TestLogger_ChangeErrorLevel(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := newLogger(buf)
	logger.ChangeErrorLevel(logrus.WarnLevel)
	logger.Error("retrying request", "request", "GET", "status", 500)
	s := buf.String()
	want := `level=warning msg="retrying request" request=GET status=500` + "\n"
	if s != want {
		t.Fatalf("wanted %s, got %s", want, s)
	}
}

func TestLogger_ChangeInfoLevel(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := newLogger(buf)
	logger.ChangeInfoLevel(logrus.WarnLevel)
	logger.Info("retrying request", "request", "GET", "status", 500)
	s := buf.String()
	want := `level=warning msg="retrying request" request=GET status=500` + "\n"
	if s != want {
		t.Fatalf("wanted %s, got %s", want, s)
	}
}

func TestLogger_ChangeWarnLevel(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := newLogger(buf)
	logger.ChangeWarnLevel(logrus.ErrorLevel)
	logger.Warn("retrying request", "request", "GET", "status", 500)
	s := buf.String()
	want := `level=error msg="retrying request" request=GET status=500` + "\n"
	if s != want {
		t.Fatalf("wanted %s, got %s", want, s)
	}
}
