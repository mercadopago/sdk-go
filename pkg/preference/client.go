package preference

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/checkout/preferences"
	urlWithID = "https://api.mercadopago.com/checkout/preferences/{id}"
	urlSearch = "https://api.mercadopago.com/checkout/preferences/search"
)

// Client contains the methods to interact with the Preference API.
type Client interface {
	// Create creates a preference with information about a product or service and obtain the URL needed to start the payment flow.
	// It is a get request to the endpoint: https://api.mercadopago.com/checkout/preferences
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/preferences/_checkout_preferences/post
	Create(ctx context.Context, dto Request) (*Response, error)

	// Get finds a preference by ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/checkout/preferences/{id}
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/preferences/_checkout_preferences_id/get
	Get(ctx context.Context, id string) (*Response, error)

	// Update updates details for a payment preference.
	// It is a get request to the endpoint: https://api.mercadopago.com/checkout/preferences/{id}
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/preferences/_checkout_preferences_id/put
	Update(ctx context.Context, id string, dto Request) (*Response, error)

	// Search finds all preference information generated through specific filters
	// It is a get request to the endpoint: https://api.mercadopago.com/checkout/preferences/search
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/preferences/_checkout_preferences_search/get
	Search(ctx context.Context, f SearchRequest) (*SearchResponsePage, error)
}

// client is the implementation of Client.
type client struct {
	config *config.Config
}

// NewClient returns a new Preference API Client.
func NewClient(c *config.Config) Client {
	return &client{
		config: c,
	}
}

func (c *client) Create(ctx context.Context, dto Request) (*Response, error) {
	body, err := json.Marshal(&dto)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlBase, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.Replace(urlWithID, "{id}", id, 1), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var formatted Response
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return &formatted, nil
}

func (c *client) Update(ctx context.Context, id string, dto Request) (*Response, error) {
	body, err := json.Marshal(&dto)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, strings.Replace(urlWithID, "{id}", id, 1), strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}

func (c *client) Search(ctx context.Context, f SearchRequest) (*SearchResponsePage, error) {
	params := url.Values{}
	params.Add("limit", strconv.Itoa(f.Limit))
	params.Add("offset", strconv.Itoa(f.Offset))
	params.Add("filters", fmt.Sprintf("%v", f.Filters))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlSearch+"?"+params.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var formatted *SearchResponsePage
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}
