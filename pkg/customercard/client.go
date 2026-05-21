// Package customercard provides a client for managing saved payment cards associated
// with customers in the MercadoPago Customer Cards API.
//
// Customer cards are tokenized payment cards stored under a customer profile.
// Managing these cards allows merchants to offer one-click checkout experiences
// by reusing previously saved card information.
//
// Use [NewClient] to create a client, then call its methods to create, retrieve,
// update, delete, or list cards for a given customer.
package customercard

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/v1/customers/{customer_id}/cards"
	urlWithID = urlBase + "/{card_id}"
)

// Client defines the interface for interacting with the MercadoPago Customer Cards API.
// It provides operations to create, retrieve, update, delete, and list saved payment
// cards for a specific customer.
type Client interface {
	// Create saves a new payment card for the customer identified by customerID.
	// The card token in the [Request] must be obtained beforehand via the Card Token API.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/cards/save-card/post
	Create(ctx context.Context, customerID string, request Request) (*Response, error)

	// Get retrieves a specific saved card identified by cardID for the customer
	// identified by customerID.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards/{card_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/cards/get-card/get
	Get(ctx context.Context, customerID, cardID string) (*Response, error)

	// Update modifies a saved card identified by cardID for the customer
	// identified by customerID. Only the fields present in the [Request] are updated.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards/{card_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/cards/update-card/put
	Update(ctx context.Context, customerID, cardID string, request Request) (*Response, error)

	// Delete removes a saved card identified by cardID from the customer
	// identified by customerID. The deleted card's data is returned in the response.
	// It is a delete request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards/{card_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/cards/delete-card/delete
	Delete(ctx context.Context, customerID, cardID string) (*Response, error)

	// List retrieves all saved payment cards for the customer identified by customerID.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/cards/get-customer-cards/get
	List(ctx context.Context, customerID string) ([]Response, error)
}

// client is the implementation of [Client].
type client struct {
	config *config.Config
}

// NewClient returns a new Customer Cards API client that uses the provided [config.Config]
// for authentication and HTTP settings.
func NewClient(c *config.Config) Client {
	return &client{
		config: c,
	}
}

func (c *client) Create(ctx context.Context, customerID string, request Request) (*Response, error) {
	pathParams := map[string]string{
		"customer_id": customerID,
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPost,
		URL:        urlBase,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, customerID, cardID string) (*Response, error) {
	pathParams := map[string]string{
		"customer_id": customerID,
		"card_id":     cardID,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Update(ctx context.Context, customerID, cardID string, request Request) (*Response, error) {
	pathParams := map[string]string{
		"customer_id": customerID,
		"card_id":     cardID,
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPut,
		URL:        urlWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Delete(ctx context.Context, customerID, cardID string) (*Response, error) {
	pathParams := map[string]string{
		"customer_id": customerID,
		"card_id":     cardID,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodDelete,
		URL:        urlWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) List(ctx context.Context, customerID string) ([]Response, error) {
	pathParams := map[string]string{
		"customer_id": customerID,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlBase,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[[]Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
