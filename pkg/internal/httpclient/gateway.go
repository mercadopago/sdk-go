package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/requester"
)

const (
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

// Send wraps needed options before send api call.
func Send(ctx context.Context, cfg *config.Config, req *http.Request) ([]byte, error) {
	req.Header.Set(productIDHeader, productID)
	req.Header.Set(acceptHeader, accept)
	req.Header.Set(contentTypeHeader, contentType)
	req.Header.Set(userAgentHeader, userAgent)
	req.Header.Set(trackingIDHeader, trackingID)
	req.Header.Set(authorizationHeader, "Bearer "+cfg.GetAccessToken())
	req.Header.Set(idempotencyHeader, uuid.New().String())

	if cfg.GetCorporationID() != "" {
		req.Header.Set(corporationIDHeader, cfg.GetCorporationID())
	}
	if cfg.GetIntegratorID() != "" {
		req.Header.Set(integratorIDHeader, cfg.GetIntegratorID())
	}
	if cfg.GetPlatformID() != "" {
		req.Header.Set(platformIDHeader, cfg.GetPlatformID())
	}

	return send(ctx, cfg.GetHTTPClient(), req)
}

func send(_ context.Context, requester requester.Requester, req *http.Request) ([]byte, error) {
	res, err := requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %w", err)
	}

	defer func() { _ = res.Body.Close() }()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &ResponseError{
			StatusCode: res.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    res.Header,
		}
	}

	if res.StatusCode > 399 {
		return nil, &ResponseError{
			StatusCode: res.StatusCode,
			Message:    string(response),
			Headers:    res.Header,
		}
	}

	return response, nil
}
