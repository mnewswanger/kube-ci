package notifiers

type notifier interface {
	fire(notificationProperties) error
	validates(notificationProperties) bool
}
