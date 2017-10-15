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

func (n *webhookNotifier) fire(m triggerMetadata) error {
	logrus.Info("Notification Firing")

	req, err := http.NewRequest(n.method, n.url, strings.NewReader(n.body))
	for header, value := range n.headers {
		req.Header.Set(header, value)
	}
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

func (n *webhookNotifier) initialize(rawProperties notificationProperties) (err error) {
	n.url = rawProperties["url"].(string)
	n.method = rawProperties["method"].(string)
	n.body = rawProperties["body"].(string)
	h := rawProperties["headers"].(map[string]interface{})
	n.headers = map[string]string{}
	for header, value := range h {
		n.headers[header] = value.(string)
	}
	return
}
