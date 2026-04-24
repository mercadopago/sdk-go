// Package merchantorder provides a client for interacting with the MercadoPago Merchant Orders API.
//
// Merchant orders are entities that group payments, shipments, and items into a single
// commercial transaction. They are automatically created when a checkout preference is
// used, or can be created manually for custom integrations. Merchant orders track the
// overall status of a transaction, including paid amounts, refunds, and shipping progress.
//
// For more information, see the MercadoPago Merchant Orders API reference:
// https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders/post
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

// Client defines the interface for interacting with the MercadoPago Merchant Orders API.
// It provides methods to create, retrieve, update, and search merchant orders.
type Client interface {
	// Get retrieves a specific merchant order by its numeric identifier.
	//
	// It performs a GET request to: https://api.mercadopago.com/merchant_orders/{id}
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders_id/get
	Get(ctx context.Context, id int) (*Response, error)

	// Search finds merchant orders matching the filters specified in [SearchRequest].
	// Results are paginated and returned as a [SearchResponse].
	//
	// It performs a GET request to: https://api.mercadopago.com/merchant_orders/search
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders_search/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Update modifies an existing merchant order identified by id with the provided [UpdateRequest] data.
	// Only the fields present in the request will be updated.
	//
	// It performs a PUT request to: https://api.mercadopago.com/merchant_orders/{id}
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders_id/put
	Update(ctx context.Context, id int, request UpdateRequest) (*Response, error)

	// Create creates a new merchant order with the provided [Request] data.
	//
	// It performs a POST request to: https://api.mercadopago.com/merchant_orders
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders/post
	Create(ctx context.Context, request Request) (*Response, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates and returns a new Merchant Orders API [Client] configured with the
// provided [config.Config]. The config must contain a valid access token for authenticating
// requests to the MercadoPago API.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Get(ctx context.Context, id int) (*Response, error) {
	pathParams := map[string]string{
		"id": strconv.Itoa(id),
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	queryParams := request.GetParams()

	requestData := httpclient.RequestData{
		Method:      http.MethodGet,
		URL:         urlSearch,
		QueryParams: queryParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Update(ctx context.Context, id int, request UpdateRequest) (*Response, error) {
	pathParams := map[string]string{
		"id": strconv.Itoa(id),
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPut,
		URL:        urlWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
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
