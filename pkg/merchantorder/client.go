package merchantorder

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/merchant_orders"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/{id}"
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
	pathParams := map[string]string{
		"id": strconv.Itoa(int(id)),
	}

	requestData := httpclient.RequestData{
		PathParams: pathParams,
		Method:     http.MethodGet,
		URL:        urlWithID,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	request.SetDefaults()

	requestData := httpclient.RequestData{
		QueryParams: request.Filters,
		Method:      http.MethodGet,
		URL:         urlSearch,
	}
	result, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Update(ctx context.Context, request UpdateRequest, id int64) (*Response, error) {
	pathParams := map[string]string{
		"id": strconv.Itoa(int(id)),
	}

	requestData := httpclient.RequestData{
		Body:       request,
		PathParams: pathParams,
		Method:     http.MethodPut,
		URL:        urlWithID,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlBase,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
