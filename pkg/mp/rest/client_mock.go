package rest

import "net/http"

type Mock struct {
	SendMock func(req *http.Request, opts ...Option) ([]byte, error)
}

func (m *Mock) Send(req *http.Request, opts ...Option) ([]byte, error) {
	return m.SendMock(req, opts...)
}
