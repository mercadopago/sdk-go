package preapproval

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/preapproval"
	urlWithID = urlBase + "/{id}"
	urlSearch = urlBase + "/search"
)

// Client contains the methods to interact with the Pre Approval API.
type Client interface {
	// Create creates a new pre-approval.
	// It is a post request to the endpoint: https://api.mercadopago.com/preapproval
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval/post
	Create(ctx context.Context, request Request) (*Response, error)

	// Get finds a pre-approval by ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/preapproval/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_id/get
	Get(ctx context.Context, id string) (*Response, error)

	// Update updates details a pre-approval by ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/preapproval/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_id/put
	Update(ctx context.Context, id string, request UpdateRequest) (*Response, error)

	// Search finds all pre-approval information generated through specific filters.
	// It is a get request to the endpoint: https://api.mercadopago.com/preapproval/search
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_search/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Pre Approval API Client.
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

func (c *client) Update(ctx context.Context, id string, request UpdateRequest) (*Response, error) {
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
