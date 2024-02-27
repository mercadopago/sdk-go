package merchantorder

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
)

const (
	urlBase   = "https://api.mercadopago.com/merchant_orders"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/:id"
)

// Client contains the methods to interact with the Merchant orders API.
type Client interface {
	// Get a specific merchant order by id.
	// It is a get request to the endpoint: https://api.mercadopago.com/merchant_orders/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders_id/get
	Get(ctx context.Context, id int64) (*Response, error)

	// Search for merchant orders.
	// It is a get request to the endpoint: https://api.mercadopago.com/merchant_orders/search
	// Reference: https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders_search/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Update a merchant order.
	// It is a put request to the endpoint: https://api.mercadopago.com/merchant_orders/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders_id/put
	Update(ctx context.Context, request UpdateRequest, id int64) (*Response, error)

	// Create a merchant order.
	// It is a post request to the endpoint: https://api.mercadopago.com/merchant_orders
	// Reference: https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders/post
	Create(ctx context.Context, request Request) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient contains the methods to interact with the merchant order API client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Get(ctx context.Context, id int64) (*Response, error) {
	param := map[string]string{
		"id": strconv.Itoa(int(id)),
	}

	result, err := baseclient.Get[*Response](ctx, c.cfg, urlWithID, baseclient.WithPathParams(param))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	params := request.Parameters()

	url, err := url.Parse(urlSearch)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}

	url.RawQuery = params

	result, err := baseclient.Get[*SearchResponse](ctx, c.cfg, url.String())

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Update(ctx context.Context, request UpdateRequest, id int64) (*Response, error) {
	param := map[string]string{
		"id": strconv.Itoa(int(id)),
	}

	result, err := baseclient.Put[*Response](ctx, c.cfg, urlWithID, request, baseclient.WithPathParams(param))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	result, err := baseclient.Post[*Response](ctx, c.cfg, urlBase, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
