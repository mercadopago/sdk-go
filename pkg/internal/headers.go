package internal

import (
	"net/http"

	"github.com/google/uuid"
)

var (
	authorizationHeader = http.CanonicalHeaderKey("authorization")
	productIDHeader     = http.CanonicalHeaderKey("x-product-id")
	idempotencyHeader   = http.CanonicalHeaderKey("x-idempotency-key")
)

func setDefaultHeaders(req *http.Request) {
	req.Header.Set(authorizationHeader, "Bearer "+_accessToken)
	req.Header.Set(productIDHeader, _productID)

	if _, ok := req.Header[idempotencyHeader]; !ok {
		req.Header.Set(idempotencyHeader, uuid.New().String())
	}
}
