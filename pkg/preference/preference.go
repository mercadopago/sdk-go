package preference

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
	"github.com/mercadopago/sdk-go/pkg/option"
)

const (
	urlBase = "https://api.mercadopago.com/checkout/preferences"
	urlWithID = "https://api.mercadopago.com/checkout/preferences/{id}"
	urlSearch = "https://api.mercadopago.com/checkout/preferences/search"
)

type Client interface {
	Create(ctx context.Context, dto Request) (*Response, error)
	Get(ctx context.Context, id string) (*Response, error)
	Update(ctx context.Context, id string, dto Request) (*Response, error)
	Search(ctx context.Context, f SearchRequest) (*SearchResponsePage, error)
}

// client is the implementation of Client.
type client struct {
	credential *credential.Credential
	config     *option.ClientOptions
}

// NewClient returns a new Preference API Client.
func NewClient(cdt *credential.Credential, opts ...option.ClientOption) Client {
	c := option.ApplyClientOptions(opts...)

	return &client{
		credential: cdt,
		config:     c,
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

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
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

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPut,  strings.Replace(urlWithID, "{id}", id, 1), strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
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

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
	if err != nil {
		return nil, err
	}

	var formatted *SearchResponsePage
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}
