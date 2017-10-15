package notifiers

import (
	"errors"
)

// Trigger represents a handle to a notification that can be run along with its execution triggers
type Trigger struct {
	Name           string            `json:"name"`
	Arguments      map[string]string `json:"arguments"`
	Events         []string          `json:"events"`
	NotifierID     string            `json:"notifier"`
	notifierHandle *Notification
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
func (t *Trigger) Fire() {
	t.notifierHandle.fire()
}
