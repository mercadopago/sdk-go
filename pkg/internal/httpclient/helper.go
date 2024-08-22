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
)

const (
	currentSDKVersion string = "1.0.4"
	productID         string = "CNITR48HSRV0CRPT3NI0"
)

var (
	pathParamRegexp = regexp.MustCompile(`{[^{}]*}`)

	userAgent  = fmt.Sprintf("MercadoPago Go SDK/%s", currentSDKVersion)
	trackingID = fmt.Sprintf("platform:%s,type:SDK%s,so;", runtime.Version(), currentSDKVersion)
)

type RequestData struct {
	Body any

	Method      string
	URL         string
	PathParams  map[string]string
	QueryParams map[string]string
}

func DoRequest[T any](ctx context.Context, cfg *config.Config, requestData RequestData) (T, error) {
	var resource T

	req, err := createRequest(ctx, cfg, requestData)
	if err != nil {
		return resource, err
	}

	b, err := Send(cfg.Requester, req)
	if err != nil {
		return resource, err
	}

	return unmarshal(b, resource)
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
		req.Header.Set("X-Idempotency-Key", uuid.New().String())
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
}

func setPathParams(req *http.Request, params map[string]string) error {
	pathURL := req.URL.Path

	for k, v := range params {
		pathParam := "{" + k + "}"
		pathURL = strings.Replace(pathURL, pathParam, v, 1)
	}

	matches := pathParamRegexp.FindAllString(pathURL, -1)
	if matches != nil {
		return fmt.Errorf("the following parameters weren't replaced: %v", matches)
	}

	req.URL.Path = pathURL
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
