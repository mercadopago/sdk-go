// Package cardtoken provides a client for creating card tokens via the MercadoPago
// Card Tokens API.
//
// Card tokens are short-lived, single-use representations of credit or debit card data.
// They allow merchants to collect card information securely without handling raw card
// numbers directly. A card token is typically created during checkout and then used
// to process a payment or to save a card to a customer profile.
//
// Use [NewClient] to create a client, then call Create to generate a new card token.
package cardtoken

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const url = "https://api.mercadopago.com/v1/card_tokens"

// Client defines the interface for interacting with the MercadoPago Card Tokens API.
type Client interface {
	// Create generates a new card token from the provided card details in the [Request].
	// The returned token is short-lived and single-use; it can be used to process a payment
	// or save a card to a customer profile.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/card_tokens
	Create(ctx context.Context, request Request) (*Response, error)
}

// client is the implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient returns a new Card Tokens API client that uses the provided [config.Config]
// for authentication and HTTP settings.
func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    url,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}
	return resource, nil
}
