package restapi

import "net/http"

//HTTPClient interface enables mocking of http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
