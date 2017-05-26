package jobs

import (
	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"
	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/rules"
)

// Job represents a KubeCI workflow
type Job struct {
	Name      string              `json:"name"`
	Notifiers []notifiers.Trigger `json:"notifiers"`
	Rules     rules.Ruleset       `json:"rules"`
	Steps     []interface{}       `json:"steps"`
}

// Trigger executes the job if it should be run
func (j *Job) Trigger(labels map[string]string) error {
	if j.shouldRun(labels) {
		panic("Job should be running!")
	}
	return nil
}

func (j *Job) shouldRun(labels map[string]string) bool {
	return true
}
