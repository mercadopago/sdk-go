package httpclient

import (
	"net/http"
)

// Mock is a test double that satisfies the [requester.Requester] interface by
// delegating to a caller-supplied function. It allows tests to control HTTP
// responses without making real network calls.
//
//	mock := &httpclient.Mock{
//	    DoMock: func(req *http.Request) (*http.Response, error) {
//	        return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{}`))}, nil
//	    },
//	}
type Mock struct {
	// DoMock is the function invoked by [Mock.Do]. Tests set this field to
	// return predetermined responses or to assert request properties.
	DoMock func(req *http.Request) (*http.Response, error)
}

// Do executes the mock function stored in [Mock.DoMock], forwarding the
// request and returning whatever the mock function produces.
func (m *Mock) Do(req *http.Request) (*http.Response, error) {
	return m.DoMock(req)
}
