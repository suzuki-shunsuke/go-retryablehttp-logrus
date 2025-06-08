package rlog_test

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-retryablehttp-logrus/rlog"
)

func Example_simple() {
	l := logrus.New()
	l.Out = os.Stdout
	l.Formatter = &logrus.TextFormatter{
		DisableTimestamp: true,
	}
	logE := logrus.NewEntry(l)
	logger := rlog.New(logE)
	logger.ChangeDebugLevel(logrus.WarnLevel)
	logger.Debug("retrying request", "request", "GET", "status", 500)
	// Output:
	// level=warning msg="retrying request" request=GET status=500
}
