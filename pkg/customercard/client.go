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

// Client contains the methods to interact with the Customer Cards API.
type Client interface {
	// Create a new customer card.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards/post
	Create(ctx context.Context, customerID string, request Request) (*Response, error)

	// Get a customer card by id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards/{card_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards_id/get
	Get(ctx context.Context, customerID, cardID string) (*Response, error)

	// Update a customer card by id.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards/{card_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards_id/put
	Update(ctx context.Context, customerID, cardID string, request Request) (*Response, error)

	// Delete deletes a customer card by id.
	// It is a delete request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards/{card_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards_id/delete
	Delete(ctx context.Context, customerID, cardID string) (*Response, error)

	// List all customer cards.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards/get
	List(ctx context.Context, customerID string) ([]Response, error)
}

// client is the implementation of Client.
type client struct {
	config *config.Config
}

// NewClient returns a new Customer Card Client.
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
		PathParams: pathParams,
		Method:     http.MethodPost,
		URL:        urlBase,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Get(ctx context.Context, customerID, cardID string) (*Response, error) {
	pathParams := map[string]string{
		"customer_id": customerID,
		"card_id":     cardID,
	}

	requestData := httpclient.RequestData{
		PathParams: pathParams,
		Method:     http.MethodGet,
		URL:        urlWithID,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Update(ctx context.Context, customerID, cardID string, request Request) (*Response, error) {
	pathParams := map[string]string{
		"customer_id": customerID,
		"card_id":     cardID,
	}

	requestData := httpclient.RequestData{
		Body:       request,
		PathParams: pathParams,
		Method:     http.MethodPut,
		URL:        urlWithID,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Delete(ctx context.Context, customerID, cardID string) (*Response, error) {
	pathParams := map[string]string{
		"customer_id": customerID,
		"card_id":     cardID,
	}

	requestData := httpclient.RequestData{
		PathParams: pathParams,
		Method:     http.MethodDelete,
		URL:        urlWithID,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) List(ctx context.Context, customerID string) ([]Response, error) {
	pathParams := map[string]string{
		"customer_id": customerID,
	}

	requestData := httpclient.RequestData{
		PathParams: pathParams,
		Method:     http.MethodGet,
		URL:        urlBase,
	}
	result, err := httpclient.DoRequest[[]Response](ctx, c.config, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
