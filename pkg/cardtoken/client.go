package cardtoken

import (
	"context"
	"fmt"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	baseURL = "https://api.mercadopago.com/v1/"
	urlGet  = baseURL + "/v1/card_tokens/{id}"
	urlPost = baseURL + "/v1/card_tokens"
)

type Client interface {
	Get(ctx context.Context, id string) (*Response, error)
	Create(ctx context.Context, request Request) (*Response, error)
}

type client struct {
	cfg *config.Config
}

func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	res, err := httpclient.Get[Response](ctx, c.cfg, strings.Replace(urlGet, "{id}", id, 1))
	if err != nil {
		return nil, fmt.Errorf("error get card token: %w", err)
	}
	return res, nil
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	res, err := httpclient.Post[Response](ctx, c.cfg, urlPost, request)
	if err != nil {
		return nil, fmt.Errorf("error create card token: %w", err)
	}
	return res, nil
}
