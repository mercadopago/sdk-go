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

// Client contains the methods to interact with the Preference API.
type Client interface {
	// Create a preference with information about a product or service and obtain the URL needed to start the payment flow.
	// It is a post request to the endpoint: https://api.mercadopago.com/checkout/preferences
	// Reference: https://www.mercadopago.com/developers/en/reference/preferences/_checkout_preferences/post
	Create(ctx context.Context, request Request) (*Response, error)

	// Get finds a preference by ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/checkout/preferences/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/preferences/_checkout_preferences_id/get
	Get(ctx context.Context, id string) (*Response, error)

	// Update updates details for a payment preference.
	// It is a put request to the endpoint: https://api.mercadopago.com/checkout/preferences/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/preferences/_checkout_preferences_id/put
	Update(ctx context.Context, id string, request Request) (*Response, error)

	// Search finds all preference information generated through specific filters
	// It is a get request to the endpoint: https://api.mercadopago.com/checkout/preferences/search
	// Reference: https://www.mercadopago.com/developers/en/reference/preferences/_checkout_preferences_search/get
	Search(ctx context.Context, request SearchRequest) (*PagingResponse, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Preference API Client.
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
		PathParams: pathParams,
		Method:     http.MethodGet,
		URL:        urlWithID,
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
		PathParams: pathParams,
		Method:     http.MethodPut,
		URL:        urlWithID,
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
		QueryParams: queryParams,
		Method:      http.MethodGet,
		URL:         urlSearch,
	}
	resource, err := httpclient.DoRequest[*PagingResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
