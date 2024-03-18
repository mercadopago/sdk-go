package requester

import "net/http"

// Requester has the minimum required method to send http requests.
type Requester interface {
	Do(req *http.Request) (*http.Response, error)
}
