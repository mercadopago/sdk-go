package defaultrequester

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func newRequestMock() *http.Request {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "", nil)

	return req
}

func newRequestMockWithBody() *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "http://test", bytes.NewBuffer([]byte(`{id:1}`)))

	return req
}

func newInvalidRequestMock() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "http://test", nil)
	req.GetBody = func() (io.ReadCloser, error) {
		return nil, fmt.Errorf("error getting body")
	}

	return req
}

func newRequestMockWithCanceledContext() *http.Request {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://test", nil)

	return req
}

func newRequestMockWithDeadlineContextAndServerError() (*http.Request, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*7))
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)

	return req, cancel
}

func newRequestWithHTTPServerUnavailableMock() (*httptest.Server, *http.Request) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))

	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, s.URL, nil)

	return s, request
}

func newRequestWithHTTPServerOKMock() (*httptest.Server, *http.Request) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// we should make this to pass in the lint pileline
		_, _ = http.ResponseWriter.Write(w, []byte(`{id:1}`))
	}))

	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, s.URL, nil)

	return s, request
}

func newRequestWithHTTPServerUnavailableAndCanceledContext() (*httptest.Server, *http.Request, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = http.ResponseWriter.Write(w, []byte(`{error}`))
	}))

	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, s.URL, nil)

	return s, request, cancel
}
