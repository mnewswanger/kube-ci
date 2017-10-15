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
	Notifiers map[string][]*notifiers.Trigger `json:"notifiers"`
	Rules     rules.Ruleset                   `json:"rules"`
	Steps     []Step                          `json:"steps"`
}

// Labels are sets of key / value pairs passed in by or derived from the reqeust
type Labels map[string]string

// Trigger executes the job if it should be run
func (j *Job) Trigger(requestLabels Labels) (err error) {
	if !j.shouldRun(requestLabels) {
		return
	}

	// Run the job in a new thread
	go func(requestLabels Labels) {
		fields := logrus.Fields{
			"job_namespace": j.Namespace,
			"job_name":      j.Name,
		}
		logrus.WithFields(fields).Info("Running Job")
		j.fireNotifiers("start")

		for _, s := range j.Steps {
			err = s.Execute(requestLabels)
			if err != nil {
				break
			}
		}

		// Handle Job Complete
		logrus.WithFields(fields).Debug("Job Complete")
		j.fireNotifiers("complete")
		if err == nil {
			logrus.WithFields(fields).Info("Job Succeeded")
			j.fireNotifiers("success")
		} else {
			// Handle Job Failure
			logrus.WithFields(fields).Info("Job Failed")
			j.fireNotifiers("failure")
		}
	}(requestLabels)
	return
}

func (j *Job) fireNotifiers(event string) {
	for _, n := range j.Notifiers[event] {
		n.Fire()
	}
}

func (j *Job) shouldRun(labels map[string]string) bool {
	return j.Rules.Matches(labels)
}
