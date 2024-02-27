package customercard

import (
	"context"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
)

const (
	urlBase   = "https://api.mercadopago.com/v1/customers/:customer_id/cards"
	urlWithID = urlBase + "/:card_id"
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
	params := map[string]string{
		"customer_id": customerID,
	}

	result, err := baseclient.Post[*Response](ctx, c.config, urlBase, request, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Get(ctx context.Context, customerID, cardID string) (*Response, error) {
	params := map[string]string{
		"customer_id": customerID,
		"card_id":     cardID,
	}

	result, err := baseclient.Get[*Response](ctx, c.config, urlWithID, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Update(ctx context.Context, customerID, cardID string, request Request) (*Response, error) {
	params := map[string]string{
		"customer_id": customerID,
		"card_id":     cardID,
	}

	result, err := baseclient.Put[*Response](ctx, c.config, urlWithID, request, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Delete(ctx context.Context, customerID, cardID string) (*Response, error) {
	params := map[string]string{
		"customer_id": customerID,
		"card_id":     cardID,
	}

	result, err := baseclient.Delete[*Response](ctx, c.config, urlWithID, nil, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) List(ctx context.Context, customerID string) ([]Response, error) {
	params := map[string]string{
		"customer_id": customerID,
	}

	result, err := baseclient.Get[[]Response](ctx, c.config, urlBase, baseclient.WithPathParams(params))
	if err != nil {
		return nil, err
	}

	return result, nil
}
