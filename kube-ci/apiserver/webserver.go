package apiserver

import (
	"strconv"
	"time"

	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"go.mikenewswanger.com/kube-ci/kube-ci/jobs"
)

var applicationIsHealthy bool
var datastore string
var configuredJobs map[string]*jobs.Job
var configuredNotifications map[string]*notifiers.Notification
var initialized bool
var logger *logrus.Logger
var verbosity uint8

// SetLogger sets the logger for the package
func SetLogger(l *logrus.Logger) {
	logger = l
}

// SetVerbosity sets the verbosity of the package
func SetVerbosity(v uint8) {
	verbosity = v
}

// StartWebserver starts an API Web Server
func StartWebserver(datastoreString string, port uint16, verbosity uint8) {
	// Prepare application for run
	t := time.Now()
	if logger == nil {
		logger = logrus.New()
	}
	logger.Level = logrus.DebugLevel
	datastore = datastoreString
	applicationIsHealthy = true

	// Configure Gin
	if verbosity == 0 {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// Set up listen endpoints
	r.GET("/healthz", getHealthz)
	r.GET("/metrics", getMetrics)

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/jobs", v1GetJobs)
			v1.GET("/notifiers", v1GetNotifiers)
			v1.GET("/reload", v1UpdateConfiguration)
			v1.POST("/hook", v1Fire)
		}
	}

	go func() {
		// Pre-cache jobs
		err := loadJobsAndNotifications()
		if err != nil {
			applicationIsHealthy = false
			panic(err)
		}

		initialized = true

		// Log the initialization
		logger.WithFields(logrus.Fields{
			"elapsed_µs": time.Since(t).Nanoseconds() / 1000,
		}).Info("Initialization complete")
	}()

	// Listen for requests
	r.Run(":" + strconv.Itoa(int(port)))
}

func getStringFromInterface(i map[string]interface{}, keys ...string) string {
	if valueRaw, exists := i[keys[0]]; exists {
		if len(keys) > 1 {
			if valueMap, ok := valueRaw.(map[string]interface{}); ok {
				return getStringFromInterface(valueMap, keys[1:]...)
			}
		} else {
			if valueString, ok := valueRaw.(string); ok {
				return valueString
			}
		}
	}
	return ""
}

func loadJobsAndNotifications() (err error) {
	loadedJobs, loadedNotifications, err := jobs.Load(datastore)
	if err != nil {
		return
	}
	configuredJobs = loadedJobs
	configuredNotifications = loadedNotifications
	return
}
