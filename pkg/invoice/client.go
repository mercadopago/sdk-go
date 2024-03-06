package invoice

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
	"net/url"
)

const (
	urlBase   = "https://api.mercadopago.com/authorized_payments"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/:id"
)

// Client contains the methods to interact with the Invoice API.
type Client interface {
	// Get finds an invoice by ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/authorized_payments/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_authorized_payments_id/get
	Get(ctx context.Context, id string) (*Response, error)

	// Search the invoices for a subscriptions by different parameters.
	// It is a get request to the endpoint: https://api.mercadopago.com/authorized_payments/search
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_authorized_payments_search/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Invoice API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	params := map[string]string{
		"id": id,
	}

	res, err := baseclient.Get[*Response](ctx, c.cfg, urlWithID, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	params := request.Parameters()

	parsedURL, err := url.Parse(urlSearch)
	if err != nil {
		return nil, fmt.Errorf("error parsing parseUrl: %w", err)
	}
	parsedURL.RawQuery = params

	res, err := baseclient.Get[*SearchResponse](ctx, c.cfg, parsedURL.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}
