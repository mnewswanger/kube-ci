package notifiers

type emailNotifier struct {
}

func (n *emailNotifier) fire(np NotificationProperties) error {
	return nil
}

func (n *emailNotifier) validates(np NotificationProperties) bool {
	return false
}
