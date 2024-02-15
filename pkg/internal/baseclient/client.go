package baseclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	currentSDKVersion string = "x.x.x"
	productID         string = "abc"
	accept            string = "application/json"
	contentType       string = "application/json; charset=UTF-8"

	headerProductID     = "X-Product-Id"
	headerAccept        = "Accept"
	headerContentType   = "Content-Type"
	headerUserAgent     = "User-Agent"
	headerTrackingID    = "X-Tracking-Id"
	headerRequestID     = "X-Request-Id"
	headerAuthorization = "Authorization"
	headerIdempotency   = "X-Idempotency-Key"

	headerCorporationID = "X-Corporation-Id"
	headerIntegratorID  = "X-Integrator-Id"
	headerPlatformID    = "X-Platform-Id"
)

var (
	userAgent  = fmt.Sprintf("MercadoPago Go SDK/%s", currentSDKVersion)
	trackingID = fmt.Sprintf("platform:%s,type:SDK%s,so;", runtime.Version(), currentSDKVersion)
)

// Get makes requests with the GET method
// Will return the struct specified in Generics
func Get[T any](ctx context.Context, cfg *config.Config, url string, opts ...Option) (T, error) {
	return make[T](ctx, cfg, url, http.MethodGet, nil, opts...)
}

// Post makes requests with the POST method
// Will return the struct specified in Generics
func Post[T any](ctx context.Context, cfg *config.Config, url string, body any, opts ...Option) (T, error) {
	return make[T](ctx, cfg, url, http.MethodPost, body, opts...)
}

// Put makes requests with the PUT method
// Will return the struct specified in Generics
func Put[T any](ctx context.Context, cfg *config.Config, url string, body any, opts ...Option) (T, error) {
	return make[T](ctx, cfg, url, http.MethodPut, body, opts...)
}

// Delete makes requests with the DELETE method
// Will return the struct specified in Generics
func Delete[T any](ctx context.Context, cfg *config.Config, url string, body any, opts ...Option) (T, error) {
	return make[T](ctx, cfg, url, http.MethodDelete, body, opts...)
}

func make[T any](ctx context.Context, cfg *config.Config, url, method string, body any, opts ...Option) (T, error) {
	var res T

	req, err := makeRequest(ctx, cfg, method, url, body, opts...)
	if err != nil {
		return res, err
	}

	b, err := httpclient.Send(cfg.Requester, req)
	if err != nil {
		return res, err
	}

	return makeResponse(b, res)
}

func makeRequest(ctx context.Context, cfg *config.Config, method, url string, body any, opts ...Option) (*http.Request, error) {
	req, err := makeHTTPRequest(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Apply all the functional options to configure the baseclient.
	opt := &clientOption{}
	for _, o := range opts {
		o(opt)
	}

	makeHeaders(req, cfg)
	makeQueryParams(req, opt.queryParams)

	if err = makePathParams(req, opt.pathParams); err != nil {
		return nil, err
	}

	return req, nil
}

func makeHTTPRequest(ctx context.Context, method, url string, body any) (*http.Request, error) {
	b, err := makeBody(body)
	if err != nil {
		return nil, err
	}

	return http.NewRequestWithContext(ctx, method, url, b)
}

func makeHeaders(req *http.Request, cfg *config.Config) {
	req.Header.Set(headerProductID, productID)
	req.Header.Set(headerAccept, accept)
	req.Header.Set(headerContentType, contentType)
	req.Header.Set(headerUserAgent, userAgent)
	req.Header.Set(headerTrackingID, trackingID)
	req.Header.Set(headerAuthorization, "Bearer "+cfg.AccessToken)
	req.Header.Set(headerIdempotency, uuid.New().String())
	req.Header.Set(headerRequestID, uuid.New().String())

	if cfg.CorporationID != "" {
		req.Header.Set(headerCorporationID, cfg.CorporationID)
	}
	if cfg.IntegratorID != "" {
		req.Header.Set(headerIntegratorID, cfg.IntegratorID)
	}
	if cfg.PlatformID != "" {
		req.Header.Set(headerPlatformID, cfg.PlatformID)
	}
}

func makeBody(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	b, err := json.Marshal(&body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	return strings.NewReader(string(b)), nil
}

func makePathParams(req *http.Request, params map[string]string) error {
	pathURL := req.URL.Path

	for k, v := range params {
		pathParam := ":" + k
		pathURL = strings.Replace(pathURL, pathParam, v, 1)
	}

	if err := validatePathParams(pathURL); err != nil {
		return err
	}

	req.URL.Path = pathURL

	return nil
}

func makeQueryParams(req *http.Request, params map[string]string) {
	queryParams := url.Values{}

	for k, v := range params {
		queryParams.Add(k, v)
	}

	req.URL.RawQuery = queryParams.Encode()
}

func makeResponse[T any](b []byte, response T) (T, error) {
	if err := json.Unmarshal(b, &response); err != nil {
		return response, err
	}

	return response, nil
}

func validatePathParams(pathURL string) error {
	if strings.Contains(pathURL, ":") {
		words := strings.Split(pathURL, "/")
		var paramsNotReplaced []string
		for _, word := range words {
			if strings.Contains(word, ":") {
				paramsNotReplaced = append(paramsNotReplaced, strings.Replace(word, ":", "", 1))
			}
		}
		return fmt.Errorf("path parameters not informed: %s", strings.Join(paramsNotReplaced, ","))
	}

	return nil
}
