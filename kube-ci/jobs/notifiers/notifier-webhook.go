package notifiers

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type webhookNotifier struct {
	cache map[string]string
}

func (n *webhookNotifier) fire(np notificationProperties) error {
	var req *http.Request
	var err error

	// // Create the HTTP request
	// switch t.Arguments["http_method"] {
	// default:
	// 	req, err = http.NewRequest("GET", t.Arguments["url"], nil)
	// }
	// if err != nil {
	// 	return err
	// }
	req, err = http.NewRequest("POST", "https://rc.clarkinc.biz/hooks/5hh6xquqfcByp47ac/wXu5GsakusR7M6k7Zo4wTH7ZFwfy9R55Qzk6gZPWotTbFzXm", strings.NewReader(`{"username":"[example-job]","text":":+1: The job succeeded}`))

	client := &http.Client{}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)

	return nil
}

func (n *webhookNotifier) validates(np notificationProperties) bool {
	// Reset the cache
	n.cache = map[string]string{}

	return true
}
