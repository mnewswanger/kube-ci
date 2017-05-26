package steps

import (
	"sync"

	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"
)

// Step represents a portion of the workflow that is run in serial
type Step struct {
	Name          string              `json:"name"`
	EventHandlers map[string]string   `json:"event_handlers"`
	Notifiers     []notifiers.Trigger `json:"notifiers"`
	Tasks         []Task              `json:"tasks"`
}

// Execute provides a method to run the step
func (s *Step) Execute() {
	var wg = &sync.WaitGroup{}
	for _, t := range s.Tasks {
		wg.Add(1)
		go func(t Task, wg *sync.WaitGroup) {
			t.run()
			wg.Done()
		}(t, wg)
	}
	wg.Wait()
}
