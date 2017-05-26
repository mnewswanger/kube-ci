package notifiers

type notifier interface {
	fire(NotificationProperties) error
	validates(NotificationProperties) bool
}
