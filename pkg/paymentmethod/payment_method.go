package paymentmethod

import (
	"encoding/json"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/mp/rest"
)

const url = "https://api.mercadopago.com/v1/payment_methods"

// Client contains the methods to interact with the Payment Methods API.
type Client interface {
	// List lists all payment methods.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payment_methods
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/payment_methods/_payment_methods/get/
	List(opts ...rest.Option) ([]Response, error)
}

// client is the implementation of Client.
type client struct {
	rc rest.Client
}

// NewClient returns a new Payment Methods API Client.
func NewClient(restClient rest.Client) Client {
	return &client{
		rc: restClient,
	}
}

func (c *client) List(opts ...rest.Option) ([]Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, &rest.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error creating request: " + err.Error(),
		}
	}

	res, err := c.rc.Send(req, opts...)
	if err != nil {
		return nil, err
	}

	var formatted []Response
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}
