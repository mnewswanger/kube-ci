package jobs

import (
	"github.com/sirupsen/logrus"
	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"
	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/rules"
)

// Job represents a KubeCI workflow
type Job struct {
	Name      string                          `json:"name"`
	Namespace string                          `json:"namespace"`
	Labels    map[string]string               `json:"labels"`
	Notifiers map[string][]*notifiers.Trigger `json:"notifiers"`
	Rules     rules.Ruleset                   `json:"rules"`
	Steps     []Step                          `json:"steps"`
}

// ShouldRun returns true when the specified job should be triggered
func (j *Job) ShouldRun(labels map[string]string) bool {
	return j.Rules.Matches(labels)
}

// Trigger executes the job if it should be run
func (j *Job) Trigger(labels map[string]string) error {
	fields := logrus.Fields{
		"job_namespace": j.Namespace,
		"job_name":      j.Name,
	}
	logrus.WithFields(fields).Info("Running Job")

	var err error
	for _, s := range j.Steps {
		err = s.Execute(labels)
		if err != nil {
			break
		}
	}

	// Handle Job Complete
	logrus.WithFields(fields).Debug("Job Complete")
	j.fireNotifiers("Complete")
	if err == nil {
		logrus.WithFields(fields).Info("Job Succeeded")
		j.fireNotifiers("success")
	} else {
		// Handle Job Failure
		logrus.WithFields(fields).Info("Job Failed")
		j.fireNotifiers("failure")
	}
	return nil
}

func (j *Job) fireNotifiers(event string) {
	for _, n := range j.Notifiers[event] {
		n.Fire()
	}
}
