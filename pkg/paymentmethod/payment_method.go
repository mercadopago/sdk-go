package paymentmethod

import (
	"context"
	"encoding/json"
	"fmt"
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
	List(ctx context.Context) ([]Response, error)
}

// client is the implementation of Client.
type client struct {
	credential *credential.Credential
	config     option.ClientOptions
}

// NewClient returns a new Payment Methods API Client.
func NewClient(cdt *credential.Credential, opts ...option.ClientOption) Client {
	options := option.ClientOptions{
		Requester: httpclient.DefaultRequester(),
	}
	for _, opt := range opts {
		opt.Apply(&options)
	}

	c := option.ClientOptions{
		Requester: options.Requester,
	}

	return &client{
		credential: cdt,
		config:     c,
	}
}

func (c *client) List(ctx context.Context) ([]Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
	if err != nil {
		return nil, err
	}

	var formatted []Response
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}
