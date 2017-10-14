package notifiers

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// Notification is an object that can interact with an external service to publish information about jobs
type Notification struct {
	Name         string                 `json:"name"`
	Namespace    string                 `json:"namespace"`
	Logger       *logrus.Logger         `json:"-"`
	Properties   notificationProperties `json:"properties"`
	Retries      uint8                  `json:"retries"`
	Type         string                 `json:"type"`
	Verbosity    uint8                  `json:"-"`
	handler      notifier
	instanceID   string
	loggerFields logrus.Fields
}

// Fire sends the specified notification
func (n Notification) fire() (err error) {
	err = n.handler.dataValidates(n.Properties)

	if err != nil {
		n.setStatus("failed")
		return
	}

	n.setStatus("pending")
	err = n.handler.fire(n.Properties)
	if err != nil {
		return
	}
	n.setStatus("succeeded")
	return
}

func (n *Notification) Register() error {
	if n.Logger == nil {
		n.Logger = logrus.New()
		switch n.Verbosity {
		case 0:
			n.Logger.Level = logrus.ErrorLevel
			break
		case 1:
			n.Logger.Level = logrus.WarnLevel
			break
		case 2:
			fallthrough
		case 3:
			n.Logger.Level = logrus.InfoLevel
			break
		default:
			n.Logger.Level = logrus.DebugLevel
			break
		}
	}

	if n.handler == nil {
		switch n.Type {
		case "webhook":
			n.handler = &webhookNotifier{
				rawProperties: n.Properties,
			}
			return n.handler.validates()
		default:
			n.Logger.WithFields(logrus.Fields{
				"name": n.Name,
				"type": n.Type,
			}).Error("Invalid type")
			return errors.New("Invalid notification type provided: " + n.Type)
		}
	}
	return nil
}

func (n Notification) setStatus(status string) {
}

type notificationProperties map[string]interface{}

type notifier interface {
	// Validate notification property data
	dataValidates(notificationProperties) error
	// Fire the notification
	fire(notificationProperties) error
	// Validate the notification properties (passed via rawProperties)
	validates() error
}
