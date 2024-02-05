package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
)

const (
	domainMP = "https://api.mercadopago.com"

	currentSDKVersion string = "x.x.x"
	productID         string = "abc"
	accept            string = "application/json"
	contentType       string = "application/json; charset=UTF-8"

	productIDHeader     = "X-Product-Id"
	acceptHeader        = "Accept"
	contentTypeHeader   = "Content-Type"
	userAgentHeader     = "User-Agent"
	trackingIDHeader    = "X-Tracking-Id"
	authorizationHeader = "Authorization"
	idempotencyHeader   = "X-Idempotency-Key"

	corporationIDHeader = "X-Corporation-Id"
	integratorIDHeader  = "X-Integrator-Id"
	platformIDHeader    = "X-Platform-Id"
)

var (
	userAgent  = fmt.Sprintf("MercadoPago Go SDK/%s", currentSDKVersion)
	trackingID = fmt.Sprintf("platform:%s,type:SDK%s,so;", runtime.Version(), currentSDKVersion)
)

// Get makes requests with the GET method
// Will return the struct specified in Generics
func Get[T any](ctx context.Context, cfg *config.Config, path string) (*T, error) {
	req, err := makeRequest(ctx, cfg, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return send[T](cfg.Requester, req)
}

// Post makes requests with the POST method
// Will return the struct specified in Generics
func Post[T any](ctx context.Context, cfg *config.Config, path string, body interface{}) (*T, error) {
	req, err := makeRequest(ctx, cfg, http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	return send[T](cfg.Requester, req)
}

// Put makes requests with the PUT method
// Will return the struct specified in Generics
func Put[T any](ctx context.Context, cfg *config.Config, path string, body interface{}) (*T, error) {
	req, err := makeRequest(ctx, cfg, http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	return send[T](cfg.Requester, req)
}

// Delete makes requests with the DELETE method
// Will return the struct specified in Generics
func Delete[T any](ctx context.Context, cfg *config.Config, path string, body interface{}) (*T, error) {
	req, err := makeRequest(ctx, cfg, http.MethodDelete, path, body)
	if err != nil {
		return nil, err
	}

	return send[T](cfg.Requester, req)
}

func makeRequest(ctx context.Context, cfg *config.Config, method, path string, body interface{}) (*http.Request, error) {
	req, err := buildHTTPRequest(ctx, method, path, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	makeHeaders(req, cfg)

	return req, nil
}

func makeHeaders(req *http.Request, cfg *config.Config) {
	req.Header.Set(productIDHeader, productID)
	req.Header.Set(acceptHeader, accept)
	req.Header.Set(contentTypeHeader, contentType)
	req.Header.Set(userAgentHeader, userAgent)
	req.Header.Set(trackingIDHeader, trackingID)
	req.Header.Set(authorizationHeader, "Bearer "+cfg.AccessToken)
	req.Header.Set(idempotencyHeader, uuid.New().String())

	if cfg.CorporationID != "" {
		req.Header.Set(corporationIDHeader, cfg.CorporationID)
	}
	if cfg.IntegratorID != "" {
		req.Header.Set(integratorIDHeader, cfg.IntegratorID)
	}
	if cfg.PlatformID != "" {
		req.Header.Set(platformIDHeader, cfg.PlatformID)
	}
}

func buildHTTPRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	b, err := buildBody(body)
	if err != nil {
		return nil, err
	}

	var url = domainMP + path

	return http.NewRequestWithContext(ctx, method, url, b)
}

func buildBody(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	b, err := json.Marshal(&body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	return strings.NewReader(string(b)), nil
}
