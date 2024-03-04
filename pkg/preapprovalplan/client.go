package preapprovalplan

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
	"net/url"
	"strings"
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
	// It is a put request to the endpoint: https://api.mercadopago.com/checkout/preferences/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan_id/put
	Update(ctx context.Context, request Request, id string) (*Response, error)

	// Search finds all pre-approval plan information generated through specific filters
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
	result, err := baseclient.Post[*Response](ctx, c.cfg, urlBase, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	result, err := baseclient.Get[*Response](ctx, c.cfg, strings.Replace(urlWithID, "{id}", id, 1))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Update(ctx context.Context, request Request, id string) (*Response, error) {
	result, err := baseclient.Put[*Response](ctx, c.cfg, strings.Replace(urlWithID, "{id}", id, 1), request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	params := request.Parameters()

	parsedUrl, err := url.Parse(urlSearch)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}
	parsedUrl.RawQuery = params

	result, err := baseclient.Get[*SearchResponse](ctx, c.cfg, parsedUrl.String())
	if err != nil {
		return nil, err
	}

	return result, nil
}
