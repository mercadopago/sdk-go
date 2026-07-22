// Package httpclient is an internal package that provides HTTP request
// construction, header management, and JSON marshalling/unmarshalling helpers
// used by the MercadoPago Go SDK's resource clients. It handles path-parameter
// substitution, query-string encoding, standard MercadoPago headers (tracking,
// idempotency, authentication), and generic response deserialisation.
package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/requestoptions"
)

const (
	currentSDKVersion string = "1.12.0"
	productID         string = "CNITR48HSRV0CRPT3NI0"
)

var (
	pathParamRegexp = regexp.MustCompile(`{[^{}]*}`)

	userAgent  = fmt.Sprintf("MercadoPago Go SDK/%s", currentSDKVersion)
	trackingID = fmt.Sprintf("platform:%s,type:SDK%s,so;", runtime.Version(), currentSDKVersion)
)

// RequestData holds all the information needed to build an HTTP request to the
// MercadoPago API. Resource clients populate this struct and pass it to
// [DoRequest] for execution.
type RequestData struct {
	// Body is the request payload that will be JSON-encoded before sending.
	// A nil Body results in a request with no body (e.g. GET, DELETE).
	Body any

	// Method is the HTTP verb (GET, POST, PUT, PATCH, DELETE, etc.).
	Method string

	// URL is the absolute endpoint URL, potentially containing path parameter
	// placeholders in the form {paramName} that will be replaced using
	// PathParams.
	URL string

	// PathParams maps placeholder names (without braces) to their substitution
	// values. For example, {"id": "123"} replaces {id} in the URL with 123.
	PathParams map[string]string

	// QueryParams maps query-string parameter names to their values, which are
	// URL-encoded and appended to the request URL.
	QueryParams map[string]string

	// IdempotencyKey overrides the auto-generated UUID in X-Idempotency-Key
	// when non-empty.
	IdempotencyKey string
}

// DoRequest builds an HTTP request from the given [RequestData], executes it
// through the [config.Config]'s [requester.Requester], and deserialises the JSON
// response body into a value of type T. It returns the zero value of T together
// with an error when the request fails at any stage (construction, transport,
// API error, or JSON unmarshalling).
//
// If ctx carries an idempotency key set via [requestoptions.WithIdempotencyKey],
// it is used as the X-Idempotency-Key header instead of an auto-generated UUID.
func DoRequest[T any](ctx context.Context, cfg *config.Config, requestData RequestData) (T, error) {
	var resource T

	if key, ok := requestoptions.IdempotencyKeyFrom(ctx); ok {
		requestData.IdempotencyKey = key
	}

	req, err := createRequest(ctx, cfg, requestData)
	if err != nil {
		return resource, err
	}

	response, err := Send(cfg.Requester, req)
	if err != nil {
		return resource, err
	}

	if len(response) == 0 {
		return resource, nil
	}

	return unmarshal(response, resource)
}

func createRequest(ctx context.Context, cfg *config.Config, requestData RequestData) (*http.Request, error) {
	body, err := marshal(requestData.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, requestData.Method, requestData.URL, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	setHeaders(req, cfg, requestData)
	if err = setPathParams(req, requestData.PathParams); err != nil {
		return nil, err
	}
	setQueryParams(req, requestData.QueryParams)

	return req, nil
}

func setHeaders(req *http.Request, cfg *config.Config, requestData RequestData) {
	req.Header.Set("X-Product-Id", productID)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Tracking-Id", trackingID)
	req.Header.Set("Authorization", "Bearer "+cfg.AccessToken)
	req.Header.Set("X-Request-Id", uuid.New().String())

	if !strings.EqualFold(requestData.Method, http.MethodGet) {
		idempotencyKey := requestData.IdempotencyKey
		if idempotencyKey == "" {
			idempotencyKey = uuid.New().String()
		}
		req.Header.Set("X-Idempotency-Key", idempotencyKey)
	}

	if cfg.CorporationID != "" {
		req.Header.Set("X-Corporation-Id", cfg.CorporationID)
	}
	if cfg.IntegratorID != "" {
		req.Header.Set("X-Integrator-Id", cfg.IntegratorID)
	}
	if cfg.PlatformID != "" {
		req.Header.Set("X-Platform-Id", cfg.PlatformID)
	}
	if cfg.ExpandNodes != "" {
		req.Header.Set("X-Expand-Responde-Nodes", cfg.ExpandNodes)
	}
}

func setPathParams(req *http.Request, params map[string]string) error {
	pathURL := req.URL.Path
	rawPathURL := req.URL.Path

	for k, v := range params {
		pathParam := "{" + k + "}"
		pathURL = strings.Replace(pathURL, pathParam, v, 1)
		rawPathURL = strings.Replace(rawPathURL, pathParam, url.PathEscape(v), 1)
	}

	matches := pathParamRegexp.FindAllString(pathURL, -1)
	if matches != nil {
		return fmt.Errorf("the following parameters weren't replaced: %v", matches)
	}

	req.URL.Path = pathURL
	req.URL.RawPath = rawPathURL
	return nil
}

func setQueryParams(req *http.Request, params map[string]string) {
	if len(params) == 0 {
		return
	}

	queryParams := url.Values{}
	for k, v := range params {
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()
}

func marshal(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	b, err := json.Marshal(&body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	return strings.NewReader(string(b)), nil
}

func unmarshal[T any](b []byte, response T) (T, error) {
	if err := json.Unmarshal(b, &response); err != nil {
		return response, err
	}

	return response, nil
}
