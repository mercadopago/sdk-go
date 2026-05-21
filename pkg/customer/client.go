// Package customer provides a client for managing customers in the MercadoPago Customers API.
//
// Customers are buyer profiles that store personal information, identification documents,
// saved cards, and addresses. Creating and managing customers simplifies checkout flows
// by enabling card reuse and pre-filled buyer data.
//
// Use [NewClient] to create a client, then call its methods to create, retrieve, update,
// or search for customers.
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

// Client defines the interface for interacting with the MercadoPago Customers API.
// It provides operations to create, retrieve, update, and search for customer records.
type Client interface {
	// Create registers a new customer with personal data, identification, and optionally saved cards
	// to simplify future payment processes.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/customers
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/customers/create-customer/post
	Create(ctx context.Context, request Request) (*Response, error)

	// Search finds customers matching the specified filters. Use [SearchRequest] to configure
	// pagination and filter criteria such as email or identification number.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customers/search
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/customers/search-customer/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Get retrieves the full profile of a customer identified by the given id, including
	// personal information, saved cards, and registered addresses.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customers/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/customers/get-customer/get
	Get(ctx context.Context, id string) (*Response, error)

	// Update modifies the data of an existing customer identified by the given id. Only the
	// fields present in the [Request] will be updated.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/customers/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/customers/update-customer/put
	Update(ctx context.Context, id string, request Request) (*Response, error)
}

// client is the implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient returns a new Customers API client that uses the provided [config.Config]
// for authentication and HTTP settings.
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

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
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

func (c *client) Update(ctx context.Context, id string, request Request) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
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
