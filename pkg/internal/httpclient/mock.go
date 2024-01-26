package httpclient

import (
	"net/http"
)

type Mock struct {
	DoMock func(req *http.Request) (*http.Response, error)
}

func (m *Mock) Do(req *http.Request) (*http.Response, error) {
	return m.DoMock(req)
}
