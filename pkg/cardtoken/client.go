// This package is useful for generate card tokens to be used on payments creation.
//
// Should be used just by PCI compliance integrators.
//
// No PCI compliance integrators should use SDK JS:
// https://github.com/mercadopago/sdk-js?tab=readme-ov-file#checkout-api.
package cardtoken

import (
	"context"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
)

const url = "https://api.mercadopago.com/v1/card_tokens"

// Client contains the method to interact with the card token API.
type Client interface {
	// Create a card token to be used on payment creation.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/card_tokens.
	Create(ctx context.Context, request Request) (*Response, error)
}

type client struct {
	cfg *config.Config
}

// NewClient returns an implementation of Client.
func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	res, err := baseclient.Post[*Response](ctx, c.cfg, url, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
