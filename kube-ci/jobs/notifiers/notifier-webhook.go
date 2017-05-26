package notifiers

type webhookNotifier struct {
}

func (n *webhookNotifier) fire(np NotificationProperties) error {
	return nil
}

func (n *webhookNotifier) validates(np NotificationProperties) bool {
	return false
}
