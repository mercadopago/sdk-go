package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/header"
	"github.com/mercadopago/sdk-go/pkg/option"
)

const (
	productID string = "abc"

	authorizationHeader = "Authorization"
	productIDHeader     = "X-Product-Id"
	idempotencyHeader   = "X-Idempotency-Key"
)

func Send(ctx context.Context, cdt credential.Credential, req *http.Request, c option.HTTPOptions) ([]byte, error) {
	for k, v := range header.Headers(ctx) {
		canonicalKey := http.CanonicalHeaderKey(k)
		req.Header[canonicalKey] = v
	}

	req.Header.Set(authorizationHeader, "Bearer "+string(cdt))
	req.Header.Set(productIDHeader, productID)
	if _, ok := req.Header[idempotencyHeader]; !ok {
		req.Header.Set(idempotencyHeader, uuid.New().String())
	}

	return send(ctx, req, c)
}

func send(ctx context.Context, req *http.Request, c option.HTTPOptions) ([]byte, error) {
	res, err := do(ctx, req, c)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %w", err)
	}

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    res.Header,
		}
	}

	if res.StatusCode > 399 {
		return nil, &ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    string(response),
			Headers:    res.Header,
		}
	}

	return response, nil
}