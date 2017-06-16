package apiserver

import (
	"github.com/sirupsen/logrus"
)

var applicationIsHealthy bool
var datastore string
var logger *logrus.Logger
