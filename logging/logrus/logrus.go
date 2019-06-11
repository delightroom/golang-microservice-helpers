package logrus

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Entry
)

func ConfigureLogger() {
	// Initialize logger to use JSONFormatter
	datadogFormatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	}
	logrus.SetFormatter(datadogFormatter)
	logrus.SetOutput(os.Stdout)

	// Create new logrusEntry
	Logger = logrus.NewEntry(logrus.StandardLogger())
	SetLogger(Logger)
}
