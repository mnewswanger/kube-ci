package jobs

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"

	"io/ioutil"

	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"
)

// Task represents a portion of the workflow that is run in serial
type Task struct {
	Name          string              `json:"name"`
	Action        string              `json:"action"`
	Arguments     map[string]string   `json:"arguments"`
	EventHandlers map[string]string   `json:"event_handlers"`
	Notifiers     []notifiers.Trigger `json:"notifiers"`
	Retries       uint8               `json:"retries"`
	Timeout       int                 `json:"timeout"`
	labels        map[string]string
}

// Run executes the task
func (t *Task) Run(labels map[string]string) error {
	t.labels = labels
	return t.run()
}

func (t *Task) run() error {
	var err error

	fields := logrus.Fields{
		"task_name": t.Name,
	}

	logrus.WithFields(fields).Info("Starting Task")

	switch t.Action {
	case "webhook":
		err = t.executeWebhook()
		break
	default:
		err = errors.New("Action type not supported (" + t.Action + ")")
	}

	logrus.WithFields(fields).Info("Task complete")
	return err
}

func (t *Task) executeWebhook() error {
	var req *http.Request
	var err error

	// Create the HTTP request
	switch t.Arguments["http_method"] {
	default:
		req, err = http.NewRequest("GET", t.Arguments["url"], nil)
	}
	if err != nil {
		return err
	}

	client := &http.Client{}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)
	return nil
}

func (t *Task) handleCompleted() {

}

func (t *Task) handleFailed() {

}

func (t *Task) handleSucceeded() {

}
