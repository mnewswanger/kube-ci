package notifiers

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type webhookNotifier struct {
	url           string
	method        string
	body          string
	headers       map[string]string
	rawProperties notificationProperties
}

func (n *webhookNotifier) fire(np notificationProperties) error {
	return nil

	var req *http.Request
	var err error

	logrus.Info("Notification Firing")

	// // Create the HTTP request
	// switch t.Arguments["http_method"] {
	// default:
	// 	req, err = http.NewRequest("GET", t.Arguments["url"], nil)
	// }
	// if err != nil {
	// 	return err
	// }
	req, err = http.NewRequest("POST", "", strings.NewReader(`{"username":"[example-job]","text":":+1: The job succeeded}`))

	client := &http.Client{}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)

	logrus.Info("Notification Fired")

	return nil
}

func (n *webhookNotifier) dataValidates(np notificationProperties) error {
	return nil
}

func (n *webhookNotifier) validates() error {
	return nil
}
