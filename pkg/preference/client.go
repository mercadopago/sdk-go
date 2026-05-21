// Package preference provides a client for interacting with the MercadoPago Checkout Preferences API.
//
// Preferences define the product or service details for a checkout session, including
// items, payer information, payment methods, shipping options, and callback URLs.
// Once created, a preference provides the init_point URL that redirects buyers to
// the MercadoPago checkout flow.
//
// For more information, see the MercadoPago Preferences API reference:
// https://www.mercadopago.com/developers/en/reference/online-payments/checkout-pro/preferences/create-preference/post
package preference

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/checkout/preferences"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/{id}"
)

// Client defines the interface for interacting with the MercadoPago Checkout Preferences API.
// It provides methods to create, retrieve, update, and search checkout preferences.
type Client interface {
	// Create creates a new checkout preference with information about a product or service
	// and returns a [Response] containing the URL needed to start the payment flow.
	//
	// It performs a POST request to: https://api.mercadopago.com/checkout/preferences
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-pro/preferences/create-preference/post
	Create(ctx context.Context, request Request) (*Response, error)

	// Get retrieves an existing checkout preference by its unique identifier.
	//
	// It performs a GET request to: https://api.mercadopago.com/checkout/preferences/{id}
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-pro/preferences/get-preference/get
	Get(ctx context.Context, id string) (*Response, error)

	// Update modifies an existing checkout preference identified by id with the provided [Request] data.
	// Only the fields present in the request will be updated.
	//
	// It performs a PUT request to: https://api.mercadopago.com/checkout/preferences/{id}
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-pro/preferences/update-preference/put
	Update(ctx context.Context, id string, request Request) (*Response, error)

	// Search finds all preferences matching the filters specified in [SearchRequest].
	// Results are paginated and returned as a [PagingResponse].
	//
	// It performs a GET request to: https://api.mercadopago.com/checkout/preferences/search
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-pro/preferences/search-preferences/get
	Search(ctx context.Context, request SearchRequest) (*PagingResponse, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates and returns a new Preference API [Client] configured with the provided [config.Config].
// The config must contain a valid access token for authenticating requests to the MercadoPago API.
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

func (c *client) Search(ctx context.Context, request SearchRequest) (*PagingResponse, error) {
	queryParams := request.GetParams()

	requestData := httpclient.RequestData{
		Method:      http.MethodGet,
		URL:         urlSearch,
		QueryParams: queryParams,
	}
	resource, err := httpclient.DoRequest[*PagingResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
