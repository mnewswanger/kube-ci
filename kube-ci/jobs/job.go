package jobs

import (
	"github.com/fatih/color"

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
	Steps     []interface{}       `json:"steps"`
}

// Trigger executes the job if it should be run
func (j *Job) Trigger(labels map[string]string) (bool, error) {
	if j.shouldRun(labels) {
		color.Green("Job (" + j.Namespace + "." + j.Name + ") should be running!")
	} else {
		color.Red("Job (" + j.Namespace + "." + j.Name + ") did not match rules")
	}
	return false, nil
}

func (j *Job) shouldRun(labels map[string]string) bool {
	return j.Rules.Matches(labels)
}
