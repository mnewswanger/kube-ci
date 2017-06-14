package jobs

import (
	"github.com/sirupsen/logrus"
	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"
	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/rules"
)

// Job represents a KubeCI workflow
type Job struct {
	Name      string              `json:"name"`
	Namespace string              `json:"namespace"`
	Labels    map[string]string   `json:"labels"`
	Notifiers []notifiers.Trigger `json:"notifiers"`
	Rules     rules.Ruleset       `json:"rules"`
	Steps     []Step              `json:"steps"`
}

// Trigger executes the job if it should be run
func (j *Job) Trigger(labels map[string]string) (bool, error) {
	if j.shouldRun(labels) {
		logrus.WithFields(logrus.Fields{
			"job_namespace": j.Namespace,
			"job_name":      j.Name,
		}).Info("Running Job")

		for _, s := range j.Steps {
			s.Execute(labels)
		}
	}
	return false, nil
}

func (j *Job) shouldRun(labels map[string]string) bool {
	return j.Rules.Matches(labels)
}
