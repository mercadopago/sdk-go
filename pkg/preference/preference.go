package preference

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
	"github.com/mercadopago/sdk-go/pkg/option"
)

const (
	url = "https://api.mercadopago.com/checkout/preferences"
	urlWithID = "https://api.mercadopago.com/checkout/preferences/{id}"
)

type Client interface {
	Create(ctx context.Context, credential credential.Credential, dto Request) (*Response, error)
	Get(ctx context.Context, credential credential.Credential, id string) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	config option.HTTPOptions
}

// NewClient returns a new Preference API Client.
func NewClient(opts ...option.HTTPOption) Client {
	options := httpclient.DefaultOptions()
	for _, opt := range opts {
		opt.ApplyHTTP(&options)
	}

	c := option.HTTPOptions{
		RetryMax:        options.RetryMax,
		HTTPClient:      options.HTTPClient,
		Timeout:         options.Timeout,
		BackoffStrategy: options.BackoffStrategy,
		CheckRetry:      options.CheckRetry,
	}
	c.HTTPClient.Timeout = c.Timeout

	return &client{config: c}
}

func (c *client) Create(ctx context.Context, cdt credential.Credential, dto Request) (*Response, error) {
	body, err := json.Marshal(&dto)
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error marshaling request body: " + err.Error(),
		}
	}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error creating request: " + err.Error(),
		}
	}

	res, err := httpclient.Send(ctx, cdt, req, c.config)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}

func (c *client) Get(ctx context.Context, cdt credential.Credential, id string) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, strings.Replace(urlWithID, "{id}", id, 1), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s" + err.Error())
	}

	res, err := httpclient.Send(ctx, cdt, req, c.config)
	if err != nil {
		return nil, err
	}

	var formatted Response
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return &formatted, nil
}
