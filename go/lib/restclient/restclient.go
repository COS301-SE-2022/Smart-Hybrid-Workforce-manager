package restclient

import "net/http"

// An HTTP client interface, used to allow easier mocking for testing

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
