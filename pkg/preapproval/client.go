package preapproval

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
	"net/url"
	"strings"
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
	Update(ctx context.Context, request UpdateRequest, id string) (*Response, error)

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

func (c *client) Update(ctx context.Context, request UpdateRequest, id string) (*Response, error) {
	result, err := baseclient.Put[*Response](ctx, c.cfg, strings.Replace(urlWithID, "{id}", id, 1), request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	params := request.Parameters()

	parsedURL, err := url.Parse(urlSearch)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}
	parsedURL.RawQuery = params

	result, err := baseclient.Get[*SearchResponse](ctx, c.cfg, parsedURL.String())
	if err != nil {
		return nil, err
	}

	return result, nil
}
