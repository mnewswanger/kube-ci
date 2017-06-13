package cmd

import (
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mikenewswanger.com/kube-ci/kube-ci/jobs"
	"go.mikenewswanger.com/utilities/filesystem"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Prepare application for run
		t := time.Now()
		logger.Level = logrus.DebugLevel
		healthy := true

		// Configure Gin
		if commandLineFlags.verbosity == 0 {
			gin.SetMode(gin.ReleaseMode)
		}
		r := gin.Default()

		// Pre-cache jobs
		configuredJobs, notifications, err := jobs.Load("filesystem", commandLineFlags.configDirectory)
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
			configuredJobs, notifications, err = jobs.Load("filesystem", commandLineFlags.configDirectory)
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

			labels["git.event"] = getStringFromInterface(d, "event_name")
			labels["git.branch"] = branch
			labels["git.tag"] = tag
			labels["git.start_commit"] = startCommit
			labels["git.target_commit"] = getStringFromInterface(d, "after")
			labels["git.project.avatar_url"] = getStringFromInterface(d, "project", "avatar_url")
			labels["git.project.name"] = getStringFromInterface(d, "project", "name")
			labels["git.project.namespace"] = getStringFromInterface(d, "project", "namespace")
			labels["git.repository.url_http"] = getStringFromInterface(d, "repository", "git_http_url")
			labels["git.repostiory.url_ssh"] = getStringFromInterface(d, "repository", "git_ssh_url")
			labels["git.user.avatar_url"] = getStringFromInterface(d, "user_avatar")
			labels["git.user.email"] = getStringFromInterface(d, "user_email")
			labels["git.user.name"] = getStringFromInterface(d, "user_name")

			ly, _ := yaml.Marshal(labels)
			color.Blue(string(ly))

			y, _ := yaml.Marshal(d)
			logger.Debug(string(y))
			c.String(200, `{"error":null,"message":"Request processed succesfully"}`)
		})

		// Listen for requests
		r.Run(":" + strconv.Itoa(int(commandLineFlags.listenPort)))

		// This is all temporary and should be moved into tests instead
		fs := filesystem.Filesystem{}
		contents, err := fs.LoadFileIfExists("~/documents/projects/kube-ci/job.yml")
		if err != nil {
			panic(err)
		}
		color.Yellow(contents)
		var job *jobs.Job
		err = yaml.Unmarshal([]byte(contents), &job)
		if err != nil {
			panic(err)
		}
		var y, _ = yaml.Marshal(job)
		color.Green(string(y))
		job.Trigger(map[string]string{})
	},
}

func init() {
	serveCmd.Flags().StringVarP(&commandLineFlags.configDirectory, "config", "c", "", "Configuration Directory")
	serveCmd.Flags().Uint16VarP(&commandLineFlags.listenPort, "listen-port", "p", 8080, "Listen Port")

	RootCmd.AddCommand(serveCmd)
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
