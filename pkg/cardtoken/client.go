package cardtoken

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	url = "https://api.mercadopago.com/v1/card_tokens"
)

type Client interface {
	Create(ctx context.Context, request Request) (*Response, error)
}

type client struct {
	cfg *config.Config
}

func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	res, err := httpclient.Post[Response](ctx, c.cfg, url, request)
	if err != nil {
		return nil, fmt.Errorf("error create card token: %w", err)
	}
	return res, nil
}
