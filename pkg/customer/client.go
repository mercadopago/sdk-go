package customer

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
)

const (
	urlBase   = "https://api.mercadopago.com/v1/customers"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/:id"
)

// Client contains the methods to interact with the Customers API.
type Client interface {
	// Create a customer with all its data and save the cards used to simplify the payment process.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/customers
	// Reference: https://www.mercadopago.com/developers/en/reference/customers/_customers/post/
	Create(ctx context.Context, request Request) (*Response, error)
	// Search find all customer information using specific filters.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customers/search
	// Reference: https://www.mercadopago.com/developers/en/reference/customers/_customers_search/get/
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
	// Get check all the information of a client created with the client ID of your choice.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customers/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/customers/_customers_id/get/
	Get(ctx context.Context, id string) (*Response, error)
	// Update renew the data of a customer.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/customers
	// Reference: https://www.mercadopago.com/developers/en/reference/customers/_customers_id/put/
	Update(ctx context.Context, id string, request Request) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Customers API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	result, err := baseclient.Post[*Response](ctx, c.cfg, urlBase, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	params := request.Parameters()

	parsedURL, err := url.Parse(urlSearch)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}
	parsedURL.RawQuery = params

	result, err := baseclient.Get[*SearchResponse](ctx, c.cfg, parsedURL.String())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	params := map[string]string{
		"id": id,
	}

	result, err := baseclient.Get[*Response](ctx, c.cfg, urlWithID, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Update(ctx context.Context, id string, request Request) (*Response, error) {
	params := map[string]string{
		"id": id,
	}

	result, err := baseclient.Put[*Response](ctx, c.cfg, urlWithID, request, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}
