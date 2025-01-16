package order

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase = "https://api.mercadopago.com/v1/orders"
)

// Client contains the methods to interact with the Order API.
type Client interface {
	// Create creates a new order.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/orders
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/order/online-payments/create/post
	Create(ctx context.Context, request Request) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Order API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlBase,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
