package paymentmethod

import (
	"encoding/json"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/api"
	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/httpclient"
)

const url = "https://api.mercadopago.com/v1/payment_methods"

// Client contains the methods to interact with the Payment Methods API.
type Client interface {
	// List lists all payment methods.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payment_methods
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/payment_methods/_payment_methods/get/
	List(credential credential.Credential, opts ...api.RequestOption) ([]Response, error)
}

// client is the implementation of Client.
type client struct {
	config httpclient.Options
}

// NewClient returns a new Payment Methods API Client.
func NewClient(opts ...httpclient.Option) Client {
	options := httpclient.DefaultOptions()

	for _, opt := range opts {
		opt.Apply(&options)
	}

	c := httpclient.Options{
		RetryMax:        options.RetryMax,
		HTTPClient:      options.HTTPClient,
		Timeout:         options.Timeout,
		BackoffStrategy: options.BackoffStrategy,
		CheckRetry:      options.CheckRetry,
	}
	c.HTTPClient.Timeout = c.Timeout

	return &client{config: c}
}

func (c *client) List(cdt credential.Credential, opts ...api.RequestOption) ([]Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error creating request: " + err.Error(),
		}
	}

	res, err := api.Send(cdt, c.config, req, opts...)
	if err != nil {
		return nil, err
	}

	var formatted []Response
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}
