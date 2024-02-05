package customercard

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	baseURL      = "https://api.mercadopago.com/v1/customers/{customer_id}"
	cardsURL     = baseURL + "/cards"
	cardsByIDURL = baseURL + cardsURL + "/{card_id}"
)

// Client contains the methods to interact with the Payment Methods API.
type Client interface {
	// Create a new customer card.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards/post
	Create(ctx context.Context, customerID string, request Request) (*Response, error)

	// Get  a customer card by ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards/{card_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards_id/get
	Get(ctx context.Context, customerID, cardID string) (*Response, error)

	// Update a customer card by ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards/{card_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards_id/put
	Update(ctx context.Context, customerID, cardID string) (*Response, error)

	// Delete deletes a customer card by ID.
	// It is a delete request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards/{card_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards_id/delete
	Delete(ctx context.Context, customerID, cardID string) (*Response, error)

	// List all customers.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/customer/{customer_id}/cards
	// Reference: https://www.mercadopago.com/developers/en/reference/cards/_customers_customer_id_cards/get
	List(ctx context.Context, customerID string) ([]Response, error)
}

// client is the implementation of Client.
type client struct {
	config *config.Config
}

// NewClient returns a new Payment Methods API Client.
func NewClient(c *config.Config) Client {
	return &client{
		config: c,
	}
}

func (c *client) Create(ctx context.Context, customerID string, request Request) (*Response, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cardsURL, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
}

func (c *client) Get(ctx context.Context, customerID, cardID string) (*Response, error) {
	url := strings.Replace(cardsByIDURL, "{customer_id}", customerID, 1)
	url = strings.Replace(cardsByIDURL, "{card_id}", cardID, 1)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
}

func (c *client) Update(ctx context.Context, customerID, cardID string) (*Response, error) {
	conv := strconv.Itoa(int(id))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.Replace(getURL, "{id}", conv, 1), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
}

func (c *client) Delete(ctx context.Context, customerID, cardID string) (*Response, error) {
	conv := strconv.Itoa(int(id))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.Replace(getURL, "{id}", conv, 1), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
}

func (c *client) List(ctx context.Context, customerID string) ([]Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var formatted []Response
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}
