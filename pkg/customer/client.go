// This package is useful for customer & cards feature and has the following features:
//   - Create customer
//   - Search customers
//   - Get customer by id
//   - Update customer
package customer

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/v1/customers"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/{id}"
)

// Client contains the methods to interact with the Customers API.
type Client interface {
	// Create a customer with all its data and can be used after to save the cards used
	// and simplify the payment process.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/customers.
	// Reference: https://www.mercadopago.com/developers/en/reference/customers/_customers/post/.
	Create(ctx context.Context, request Request) (*Response, error)

	// Search all customers with the sent filters.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customers/search.
	// Reference: https://www.mercadopago.com/developers/en/reference/customers/_customers_search/get/.
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Get return all customer data with the sent id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customers/{id}.
	// Reference: https://www.mercadopago.com/developers/en/reference/customers/_customers_id/get/.
	Get(ctx context.Context, id string) (*Response, error)

	// Update customer data with the sent id.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/customers.
	// Reference: https://www.mercadopago.com/developers/en/reference/customers/_customers_id/put/.
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

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	queryParams := request.GetParams()

	requestData := httpclient.RequestData{
		QueryParams: queryParams,
		Method:      http.MethodGet,
		URL:         urlSearch,
	}
	result, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
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

func (c *client) Update(ctx context.Context, id string, request Request) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
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
