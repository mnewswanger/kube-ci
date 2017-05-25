package steps

import (
	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"
)

// Step represents a portion of the workflow that is run in serial
type Step struct {
	Name          string              `json:"name"`
	EventHandlers map[string]string   `json:"event_handlers"`
	Notifiers     []notifiers.Trigger `json:"notifiers"`
	Tasks         []Task              `json:"tasks"`
}
