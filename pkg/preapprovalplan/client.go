package preapprovalplan

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/preapproval_plan"
	urlWithID = urlBase + "/{id}"
	urlSearch = urlBase + "/search"
)

// Client contains the methods to interact with the Pre Approval Plan API.
type Client interface {
	// Create creates a new pre-approval plan.
	// It is a post request to the endpoint: https://api.mercadopago.com/preapproval_plan
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan/post/
	Create(ctx context.Context, request Request) (*Response, error)

	// Get finds a pre-approval plan by ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/preapproval_plan/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan_id/get
	Get(ctx context.Context, id string) (*Response, error)

	// Update updates details a pre-approval plan by ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/preapproval_plan/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan_id/put
	Update(ctx context.Context, id string, request Request) (*Response, error)

	// Search finds all pre-approval plan information generated through specific filters.
	// It is a get request to the endpoint: https://api.mercadopago.com/preapproval_plan/search
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan_search/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Pre Approval Plan API Client.
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

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	queryParams := request.GetParams()

	requestData := httpclient.RequestData{
		QueryParams: queryParams,
		Method:      http.MethodGet,
		URL:         urlSearch,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
