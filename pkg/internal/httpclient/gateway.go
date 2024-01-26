package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/header"
	"github.com/mercadopago/sdk-go/pkg/option"
)

const (
	productID string = "abc"

	authorizationHeader = "Authorization"
	productIDHeader     = "X-Product-Id"
	idempotencyHeader   = "X-Idempotency-Key"
)

func Send(ctx context.Context, c *config.Config, req *http.Request) ([]byte, error) {
	for k, v := range header.Headers(ctx) {
		canonicalKey := http.CanonicalHeaderKey(k)
		req.Header[canonicalKey] = v
	}

	req.Header.Set(authorizationHeader, "Bearer "+c.AccessToken)
	req.Header.Set(productIDHeader, productID)
	if _, ok := req.Header[idempotencyHeader]; !ok {
		req.Header.Set(idempotencyHeader, uuid.New().String())
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
