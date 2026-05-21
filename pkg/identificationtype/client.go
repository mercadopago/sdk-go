// Package identificationtype provides a client for retrieving accepted identification
// document types from the MercadoPago Identification Types API.
//
// Identification types represent the kinds of identity documents (e.g., CPF, DNI, CC, RUT)
// accepted in each country where MercadoPago operates. They are typically used to validate
// buyer or cardholder identification data during payment or customer registration flows.
//
// Use [NewClient] to create a client, then call List to retrieve all available document types.
package identificationtype

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const url = "https://api.mercadopago.com/v1/identification_types"

// Client defines the interface for interacting with the MercadoPago Identification Types API.
type Client interface {
	// List retrieves all identification document types available for the country associated
	// with the authenticated credentials. Each entry includes the document type ID, name,
	// and allowed length constraints.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/identification_types
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/identification-types/get
	List(ctx context.Context) ([]Response, error)
}

// client is the implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient returns a new Identification Types API client that uses the provided
// [config.Config] for authentication and HTTP settings.
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
