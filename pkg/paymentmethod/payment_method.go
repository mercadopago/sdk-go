package paymentmethod

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
	"github.com/mercadopago/sdk-go/pkg/option"
)

const url = "https://api.mercadopago.com/v1/payment_methods"

// Client contains the methods to interact with the Payment Methods API.
type Client interface {
	// List lists all payment methods.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payment_methods
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/payment_methods/_payment_methods/get/
	List(ctx context.Context, credential credential.Credential) ([]Response, error)
}

// client is the implementation of Client.
type client struct {
	config option.HTTPOptions
}

// NewClient returns a new Payment Methods API Client.
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

func (c *client) List(ctx context.Context, cdt credential.Credential) ([]Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

	var formatted []Response
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}
