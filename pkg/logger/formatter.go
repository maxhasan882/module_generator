package logger

import "github.com/sirupsen/logrus"

func GetLogger(module string) *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	log := logger.WithFields(logrus.Fields{
		"module": module,
	})
	return log
}
