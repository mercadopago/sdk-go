// Package preapprovalplan provides a client for the MercadoPago Subscription Plan
// (Pre-Approval Plan) API.
//
// A pre-approval plan is a seller-defined subscription template that specifies recurring
// billing rules such as amount, frequency, currency, and allowed payment methods. Payers
// subscribe to a plan through the [preapproval] package, which creates individual
// subscriptions linked to these templates.
//
// Use this package to create, retrieve, update, and search subscription plan templates.
//
// API reference: https://www.mercadopago.com/developers/en/reference/subscriptions
package preapprovalplan

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	// urlBase is the base endpoint for the Pre-Approval Plan API.
	urlBase = "https://api.mercadopago.com/preapproval_plan"
	// urlWithID is the endpoint for operating on a specific pre-approval plan by its ID.
	urlWithID = urlBase + "/{id}"
	// urlSearch is the endpoint for searching pre-approval plans with filters.
	urlSearch = urlBase + "/search"
)

// Client provides methods to interact with the MercadoPago Pre-Approval Plan
// (Subscription Plan) API. It supports creating, retrieving, updating, and
// searching subscription plan templates.
type Client interface {
	// Create creates a new pre-approval plan (subscription template) for the seller.
	// The [Request] body must include recurrence settings and optionally allowed payment methods.
	// It is a POST request to the endpoint: https://api.mercadopago.com/preapproval_plan
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan/post/
	Create(ctx context.Context, request Request) (*Response, error)

	// Get retrieves a pre-approval plan (subscription template) by its unique identifier.
	// It is a GET request to the endpoint: https://api.mercadopago.com/preapproval_plan/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan_id/get
	Get(ctx context.Context, id string) (*Response, error)

	// Update modifies an existing pre-approval plan (subscription template) identified by id.
	// The same [Request] struct is used; only non-zero fields are sent to the API.
	// It is a PUT request to the endpoint: https://api.mercadopago.com/preapproval_plan/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan_id/put
	Update(ctx context.Context, id string, request Request) (*Response, error)

	// Search finds pre-approval plans (subscription templates) that match the filters
	// specified in [SearchRequest]. Results are paginated; use Limit and Offset to control pagination.
	// It is a GET request to the endpoint: https://api.mercadopago.com/preapproval_plan/search
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan_search/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
}

// client is the internal implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates a new Pre-Approval Plan API client using the provided [config.Config].
// The configuration must include valid authentication credentials.
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
