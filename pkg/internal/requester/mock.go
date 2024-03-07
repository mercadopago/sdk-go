package requester

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func NewRequestMock() *http.Request {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://test", nil)

	return req
}

func NewRequestMockWithBody() *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "http://test", bytes.NewBuffer([]byte(`{id:1}`)))

	return req
}

func NewInvalidRequestMock() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "http://test", nil)
	req.GetBody = func() (io.ReadCloser, error) {
		return nil, fmt.Errorf("error getting body")
	}

	return req
}

func NewRequestMockWithCancelledContext() *http.Request {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://test", nil)

	return req
}

func NewRequestMockWithDeadlineContext() (*http.Request, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*7))
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://test", nil)

	return req, cancel
}

func NewRequestWithHTTPServerUnavailableMock() (*httptest.Server, *http.Request) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))

	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, s.URL, nil)

	return s, request
}

func NewRequestWithHTTPServerOKMock() (*httptest.Server, *http.Request) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{id:1}`))
	}))

	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, s.URL, nil)

	return s, request
}
