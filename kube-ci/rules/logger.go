package rules

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Level = logrus.ErrorLevel
	logger.Level = logrus.DebugLevel
}
