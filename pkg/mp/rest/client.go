package rest

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	productID = "123"

	defaultTimeout = time.Duration(time.Second * 30)
)

var (
	authorizationHeader = http.CanonicalHeaderKey("authorization")
	productIDHeader     = http.CanonicalHeaderKey("x-product-id")
	idempotencyHeader   = http.CanonicalHeaderKey("x-idempotency-key")
)

var c *client

// Client is the interface that wraps the basic Send method.
type Client interface {

	// Send sends a request to the API.
	// opts are optional parameters to be used in the request, if you do not need, ignore it.
	Send(req *http.Request, opts ...Option) ([]byte, error)
}

// client is the implementation of Client.
type client struct {
	accessToken string
	productID   string

	httpClient  *http.Client
	retryClient RetryClient
}

func NewClient(at string) Client {
	c = &client{
		accessToken: at,
		productID:   productID,
		httpClient:  &http.Client{},
		retryClient: &retryClient{},
	}
	return c
}

func SetAT(at string) {
	c.accessToken = at
}

func SetHC(hc *http.Client) {
	c.httpClient = hc
}

func SetRC(rc RetryClient) {
	c.retryClient = rc
}

func (cl *client) Send(req *http.Request, opts ...Option) ([]byte, error) {
	cl.prepareRequest(req, opts...)

	res, err := c.httpClient.Do(req)
	if shouldRetry(res, err) {
		res, err = c.retryClient.Retry(req, c.httpClient, opts...)
	}
	if err != nil {
		return nil, &ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    "error sending request: " + err.Error(),
		}
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    res.Header,
		}
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, &ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    string(response),
			Headers:    res.Header,
		}
	}

	return response, nil
}

func (cl *client) prepareRequest(req *http.Request, opts ...Option) {
	timeout := defaultTimeout

	options := &options{}
	for _, opt := range opts {
		opt.apply(options)
	}
	if options.timeout > 0 {
		timeout = options.timeout
	}
	ctx, cancel := context.WithTimeout(req.Context(), timeout)
	defer cancel()
	req = req.WithContext(ctx)
	if options.customHeaders != nil {
		for k, v := range options.customHeaders {
			canonicalKey := http.CanonicalHeaderKey(k)
			req.Header[canonicalKey] = v
		}
	}
	setDefaultHeaders(req)
}

func setDefaultHeaders(req *http.Request) {
	req.Header.Add(authorizationHeader, "Bearer "+c.accessToken)
	req.Header.Add(productIDHeader, c.productID)

	if _, ok := req.Header[idempotencyHeader]; !ok {
		req.Header.Add(idempotencyHeader, uuid.New().String())
	}
}

func shouldRetry(res *http.Response, err error) bool {
	return err != nil || res.StatusCode >= http.StatusInternalServerError
}
