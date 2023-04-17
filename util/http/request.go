package http

import (
	"net/http"
)

// MakeHTTPRequest makes a HTTP request to the given URL with the given headers
func MakeHTTPRequest(url string, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest("POST", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	return resp, err
}
