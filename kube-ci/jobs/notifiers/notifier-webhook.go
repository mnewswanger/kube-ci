package notifiers

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/sirupsen/logrus"
)

type webhookNotifier struct {
	url                    *template.Template
	method                 string
	body                   *template.Template
	headers                map[string]*template.Template
	disableSSLVerification bool
	rawProperties          notificationProperties
}

func (n *webhookNotifier) fire(m triggerMetadata) error {
	logrus.Info("Notification Firing")

	url := bytes.Buffer{}
	writer := bufio.NewWriter(&url)
	n.url.Execute(writer, m)
	writer.Flush()

	body := bytes.Buffer{}
	writer = bufio.NewWriter(&body)
	n.body.Execute(writer, m)
	writer.Flush()

	req, err := http.NewRequest(n.method, url.String(), bufio.NewReader(&body))
	for header, value := range n.headers {
		v := bytes.Buffer{}
		writer = bufio.NewWriter(&v)
		value.Execute(writer, m)
		writer.Flush()
		req.Header.Set(header, v.String())
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: n.disableSSLVerification},
		},
	}

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
	n.url, err = template.New("").Parse(rawProperties["url"].(string))
	if err != nil {
		return
	}
	n.method = rawProperties["method"].(string)
	n.body, err = template.New("").Parse(rawProperties["body"].(string))
	if err != nil {
		return
	}
	h := rawProperties["headers"].(map[string]interface{})
	n.headers = map[string]*template.Template{}
	for header, value := range h {
		t, err := template.New("").Parse(value.(string))
		if err != nil {
			return err
		}
		n.headers[header] = t
	}
	disableSSLVerification, exists := rawProperties["disableSSLVerification"]
	if exists {
		n.disableSSLVerification = disableSSLVerification.(bool)
	}
	return
}
