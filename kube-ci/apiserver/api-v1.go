package apiserver

import (
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"go.mikenewswanger.com/kube-ci/kube-ci/jobs"
)

var v1Fire = func(c *gin.Context) {
	d := map[string]interface{}{}
	err := c.Bind(&d)

	logger.Debug(d)
	if err != nil {
		logger.Error(err)
	}
	var labels = jobs.Labels{}

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

	for _, j := range configuredJobs {
		j.Trigger(labels)
	}

	c.String(200, `{"error":null,"message":"Request processed succesfully"}`)
}

var v1GetJobs = func(c *gin.Context) {
	c.JSON(200, configuredJobs)
}

var v1GetNotifiers = func(c *gin.Context) {
	c.JSON(200, configuredNotifications)
}

var v1UpdateConfiguration = func(c *gin.Context) {
	err := loadJobsAndNotifications()
	if err != nil {
		c.String(500, `{"error": "Could not reload configuration"}`)
	}
	c.String(200, `{"error": null}`)
}
