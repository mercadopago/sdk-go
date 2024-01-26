package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/option"
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
	userAgent  string = fmt.Sprintf("MercadoPago Go SDK/%s", currentSDKVersion)
	trackingID string = fmt.Sprintf("platform:%s,type:SDK%s,so;", runtime.Version(), currentSDKVersion)
)

func Send(ctx context.Context, c *config.Config, req *http.Request) ([]byte, error) {
	req.Header.Set(productIDHeader, productID)
	req.Header.Set(acceptHeader, accept)
	req.Header.Set(contentTypeHeader, contentType)
	req.Header.Set(userAgentHeader, userAgent)
	req.Header.Set(trackingIDHeader, trackingID)
	req.Header.Set(authorizationHeader, "Bearer "+c.AccessToken)
	req.Header.Set(idempotencyHeader, uuid.New().String())

	if c.CorporationID != "" {
		req.Header.Set(corporationIDHeader, c.CorporationID)
	}
	if c.IntegratorID != "" {
		req.Header.Set(integratorIDHeader, c.IntegratorID)
	}
	if c.PlatformID != "" {
		req.Header.Set(platformIDHeader, c.PlatformID)
	}

	return send(ctx, c.HTTPClient, req)
}

func send(ctx context.Context, requester option.Requester, req *http.Request) ([]byte, error) {
	res, err := requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %w", err)
	}

	response, err := io.ReadAll(res.Body)
	defer res.Body.Close()
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
