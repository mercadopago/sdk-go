// Package preapproval provides a client for the MercadoPago Subscriptions (Pre-Approval) API.
//
// A pre-approval represents an individual subscription created by a payer, optionally linked
// to a [preapprovalplan]. It manages the lifecycle of recurring charges including creation,
// status updates, and search operations.
//
// For subscription plan templates (seller-side), see the [preapprovalplan] package instead.
//
// API reference: https://www.mercadopago.com/developers/en/reference/online-payments/subscriptions/create-preapproval/post
package preapproval

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	// urlBase is the base endpoint for the Pre-Approval API.
	urlBase = "https://api.mercadopago.com/preapproval"
	// urlWithID is the endpoint for operating on a specific pre-approval by its ID.
	urlWithID = urlBase + "/{id}"
	// urlSearch is the endpoint for searching pre-approvals with filters.
	urlSearch = urlBase + "/search"
)

// Client provides methods to interact with the MercadoPago Pre-Approval (Subscriptions) API.
// It supports creating, retrieving, updating, and searching subscriptions.
type Client interface {
	// Create creates a new pre-approval (subscription) for a payer.
	// The [Request] body must include payer information and recurrence settings.
	// It is a POST request to the endpoint: https://api.mercadopago.com/preapproval
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/subscriptions/create-preapproval/post
	Create(ctx context.Context, request Request) (*Response, error)

	// Get retrieves a pre-approval (subscription) by its unique identifier.
	// It is a GET request to the endpoint: https://api.mercadopago.com/preapproval/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/subscriptions/get-preapproval/get
	Get(ctx context.Context, id string) (*Response, error)

	// Update modifies an existing pre-approval (subscription) identified by id.
	// Only the fields present in [UpdateRequest] will be changed; absent fields are left unchanged.
	// It is a PUT request to the endpoint: https://api.mercadopago.com/preapproval/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/subscriptions/update-preapproval/put
	Update(ctx context.Context, id string, request UpdateRequest) (*Response, error)

	// Search finds pre-approvals (subscriptions) that match the filters specified in [SearchRequest].
	// Results are paginated; use Limit and Offset to control pagination.
	// It is a GET request to the endpoint: https://api.mercadopago.com/preapproval/search
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/subscriptions/search-preapproval/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
}

// client is the internal implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates a new Pre-Approval API client using the provided [config.Config].
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

func (c *client) Update(ctx context.Context, id string, request UpdateRequest) (*Response, error) {
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
