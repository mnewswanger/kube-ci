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
func (n Notification) fire() error {
	var err error
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
			n.handler = &webhookNotifier{}
			break
		default:
			n.Logger.WithFields(logrus.Fields{
				"name": n.Name,
				"type": n.Type,
			}).Error("Invalid type")
			err = errors.New("Invalid notification type provided: " + n.Type)
		}
	}

	if err == nil {
		if n.handler.validates(n.Properties) {
			n.setStatus("pending")
			err = n.handler.fire(n.Properties)
			if err == nil {
				n.setStatus("succeeded")
				return err
			}
		} else {
			err = errors.New("Notification validation failed")
		}
	}

	n.setStatus("failed")
	return err
}

func (n Notification) setStatus(status string) {
}

type notificationProperties map[string]interface{}
