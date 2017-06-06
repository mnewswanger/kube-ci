package notifiers

// Trigger represents a handle to a notification that can be run along with its execution triggers
type Trigger struct {
	Name      string            `json:"name"`
	Arguments map[string]string `json:"arguments"`
	Events    string            `json:"events"`
	Notifier  string            `json:"notifier"`
}
