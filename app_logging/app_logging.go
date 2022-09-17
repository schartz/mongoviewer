package app_logging

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var AppLogger *logrus.Logger

func GetLogger() *logrus.Logger {
	var initOnce sync.Once
	initOnce.Do(func() {
		AppLogger = logrus.New()
	})
	return AppLogger
}
