package apiserver

import (
	"strconv"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"go.mikenewswanger.com/kube-ci/kube-ci/jobs"
)

// StartWebserver starts an API Web Server
func StartWebserver(datastore string, port uint16, verbosity uint8) {
	// Prepare application for run
	t := time.Now()
	if logger == nil {
		logger = logrus.New()
	}
	logger.Level = logrus.DebugLevel
	healthy := true

	// Configure Gin
	if verbosity == 0 {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// Pre-cache jobs
	configuredJobs, notifications, err := jobs.Load(datastore)
	if err != nil {
		panic(err)
	}

	// Log the initialization
	logger.WithFields(logrus.Fields{
		"elapsed_µs": time.Since(t).Nanoseconds() / 1000,
	}).Info("Initialization complete")

	// Set up listen endpoints
	r.GET("/healthz", func(c *gin.Context) {
		if healthy {
			c.String(200, `{"status": "up"}`)
		} else {
			c.String(500, `{"status": "down"}`)
		}
	})
	r.GET("/jobs", func(c *gin.Context) {
		c.JSON(200, configuredJobs)
	})
	r.GET("/notifiers", func(c *gin.Context) {
		c.JSON(200, notifications)
	})
	r.GET("/reload", func(c *gin.Context) {
		configuredJobs, notifications, err = jobs.Load(datastore)
		if err != nil {
			c.String(500, `{"error": "Could not reload configuration"}`)
		}
		c.String(200, `{"error": null}`)
	})

	r.POST("/", func(c *gin.Context) {
		d := map[string]interface{}{}
		err := c.Bind(&d)
		if err != nil {
			logger.Error(err)
		}
		var labels = map[string]string{}

		// Attempt to add git labels from gitlab webhook structure
		ref := getStringFromInterface(d, "ref")
		tag := strings.TrimPrefix(ref, "refs/tags/")
		if tag == ref {
			tag = ""
		}
		branch := strings.TrimPrefix(ref, "refs/heads/")
		if branch == ref {
			branch = ""
		}
		startCommit := getStringFromInterface(d, "before")
		if startCommit == "0000000000000000000000000000000000000000" {
			startCommit = ""
		}
		targetCommit := getStringFromInterface(d, "after")
		if targetCommit == "0000000000000000000000000000000000000000" {
			targetCommit = ""
		}

		labels["git.event"] = getStringFromInterface(d, "event_name")
		labels["git.branch"] = branch
		labels["git.tag"] = tag
		labels["git.start_commit"] = startCommit
		labels["git.target_commit"] = targetCommit
		labels["git.project.avatar_url"] = getStringFromInterface(d, "project", "avatar_url")
		labels["git.project.name"] = getStringFromInterface(d, "project", "name")
		labels["git.project.namespace"] = getStringFromInterface(d, "project", "namespace")
		labels["git.repository.url_http"] = getStringFromInterface(d, "repository", "git_http_url")
		labels["git.repository.url_ssh"] = getStringFromInterface(d, "repository", "git_ssh_url")
		labels["git.user.avatar_url"] = getStringFromInterface(d, "user_avatar")
		labels["git.user.email"] = getStringFromInterface(d, "user_email")
		labels["git.user.name"] = getStringFromInterface(d, "user_name")

		switch verbosity {
		case 5:
			y, _ := yaml.Marshal(d)
			logger.Debug(string(y))
			fallthrough
		case 4:
			labelFields := logrus.Fields{}
			for k, v := range labels {
				labelFields[k] = v
			}
			logger.WithFields(labelFields).Info("Processed Labels")
		}

		triggeredJobs := []struct {
			namespace string
			name      string
		}{}
		for _, j := range configuredJobs {
			running, err := j.Trigger(labels)
			if err != nil {
				logger.Error(err)
			}
			if running {
				triggeredJobs = append(triggeredJobs, struct {
					namespace string
					name      string
				}{
					namespace: j.Namespace,
					name:      j.Name,
				})
			}
		}

		c.String(200, `{"error":null,"message":"Request processed succesfully"}`)
	})

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
