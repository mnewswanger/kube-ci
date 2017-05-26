package notifiers

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// Notification is an object that can interact with an external service to publish information about jobs
type Notification struct {
	Name         string                 `json:"name"`
	Logger       *logrus.Logger         `json:"-"`
	Properties   NotificationProperties `json:"properties"`
	Type         string                 `json:"notification_type"`
	Verbosity    uint8                  `json:"-"`
	handler      notifier
	loggerFields logrus.Fields
}

// Fire sends the specified notification
func (n Notification) Fire() error {
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
		case "email":
			n.handler = &emailNotifier{}
			break
		case "webhook":
			n.handler = &webhookNotifier{}
			break
		default:
			n.Logger.WithFields(logrus.Fields{
				"name": n.Name,
				"type": n.Type,
			}).Error("Invalid type")
			return errors.New("Invalid notification type provided: " + n.Type)
		}
	}
	if !n.handler.validates(n.Properties) {
		return errors.New("Notification validation failed")
	}
	return n.handler.fire(n.Properties)
}

// NotificationProperties represent metadata that can be sent to a notification
type NotificationProperties map[string]interface{}
