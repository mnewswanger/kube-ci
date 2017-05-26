package steps

import (
	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"
)

// Task represents a portion of the workflow that is run in serial
type Task struct {
	Name          string              `json:"name"`
	Action        string              `json:"string"`
	Arguments     map[string]string   `json:"arguments"`
	EventHandlers map[string]string   `json:"event_handlers"`
	Notifiers     []notifiers.Trigger `json:"notifiers"`
	Retries       uint8               `json:"retries"`
}

func (t *Task) run() {

}
