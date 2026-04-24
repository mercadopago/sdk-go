// Package paymentmethod provides a client for the MercadoPago Payment Methods API.
//
// It allows listing the payment methods available for the authenticated account,
// including card brands, bank transfers, and offline methods.
// Use [NewClient] to obtain a [Client] and interact with the API.
//
// For full API documentation see https://www.mercadopago.com/developers/en/reference/payment_methods/_payment_methods/get/.
package paymentmethod

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const url = "https://api.mercadopago.com/v1/payment_methods"

// Client provides methods to interact with the MercadoPago Payment Methods API.
// Create one via [NewClient].
type Client interface {
	// List retrieves all payment methods available for the authenticated account.
	// The returned slice contains details such as accepted card brands, minimum and
	// maximum amounts, accreditation times, and required additional information.
	//
	// GET https://api.mercadopago.com/v1/payment_methods
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/payment_methods/_payment_methods/get/
	List(ctx context.Context) ([]Response, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates a new Payment Methods API client using the provided [config.Config].
// The configuration must include valid credentials for authenticating with the MercadoPago API.
func NewClient(cfg *config.Config) Client {
	return &client{cfg}
}

func (c *client) List(ctx context.Context) ([]Response, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    url,
	}
	resource, err := httpclient.DoRequest[[]Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
