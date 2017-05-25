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
