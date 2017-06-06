package cmd

import (
	"github.com/sirupsen/logrus"
)

type flags struct {
	configDirectory string
	listenPort      uint16
	verbosity       int
}

var commandLineFlags = flags{}

var logger = logrus.New()
