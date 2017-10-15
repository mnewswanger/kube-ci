package notifiers

import (
	"errors"
)

// Trigger represents a handle to a notification that can be run along with its execution triggers
type Trigger struct {
	Name           string                 `json:"name"`
	Properties     map[string]interface{} `json:"properties"`
	Events         []string               `json:"events"`
	NotifierID     string                 `json:"notifier"`
	notifierHandle *Notification
}

type JobProperties struct {
	Event         string
	Name          string
	Namespace     string
	RequestLabels map[string]string
}

type triggerMetadata struct {
	job          JobProperties
	notification map[string]interface{}
}

// Bind binds the trigger to a notifier
func (t *Trigger) Bind(n map[string]*Notification) error {
	var exists bool
	t.notifierHandle, exists = n[t.NotifierID]
	if !exists {
		return errors.New("Could not bind notifier: " + t.NotifierID)
	}
	return nil
}

// Fire executes a trigger
func (t *Trigger) Fire(jobMetadata JobProperties) {
	metadata := triggerMetadata{
		job:          jobMetadata,
		notification: t.Properties,
	}
	t.notifierHandle.fire(metadata)
}
