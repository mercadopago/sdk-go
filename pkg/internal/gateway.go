package internal

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/httpclient"
)

const (
	productID string = "abc"

	authorizationHeader = "Authorization"
	productIDHeader     = "X-Product-Id"
	idempotencyHeader   = "X-Idempotency-Key"
)

func Send(cdt credential.Credential, requester httpclient.Requester, req *http.Request) ([]byte, error) {
	req.Header.Set(authorizationHeader, "Bearer "+string(cdt))
	req.Header.Set(productIDHeader, productID)
	if _, ok := req.Header[idempotencyHeader]; !ok {
		req.Header.Set(idempotencyHeader, uuid.New().String())
	}

	return send(requester, req)
}

func send(requester httpclient.Requester, req *http.Request) ([]byte, error) {
	res, err := requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %s", err.Error())
	}

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    res.Header,
		}
	}

	if res.StatusCode > 399 {
		return nil, &httpclient.ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    string(response),
			Headers:    res.Header,
		}
	}

	return response, nil
}
