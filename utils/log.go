package utils

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.DebugLevel
}

func PrintLog(level string, message string) {
	switch level {
		case "debug":
			log.Debug(message)
		case "info":
			log.Info(message)
		case "warn":
			log.Warn(message)
		case "error":
			log.Error(message)
	}
}