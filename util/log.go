package util

import (
	"time"

	"github.com/leogaonabr/golang-service/config"
	"github.com/sirupsen/logrus"
)

var baseLogger *logrus.Logger

// InitLogger creates the template logger with the default fields
func InitLogger() {
	baseLogger = logrus.New()
	baseLogger.SetNoLock()

	baseLogger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	}

}

// GetLogger returns the template logger instance
func GetLogger() *logrus.Entry {
	if baseLogger == nil {
		InitLogger()
	}

	return baseLogger.WithFields(map[string]interface{}{
		"app":         "golang-template-app",
		"version":     config.GetVersion(),
		"environment": config.GetEnv(),
	})
}
