package refund

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	baseURL = "https://api.mercadopago.com/v1/payments/{id}/"
	getURL  = baseURL + "refunds/{refund_id}"
	listURL = baseURL + "refunds"
	postURL = baseURL + "refunds"
)

// Client contains the methods to interact with the Payment's refund API.
type Client interface {
	// Get gets a specific refund by payment id and refund id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds/{refund_id}
	// Reference: https://www.mercadopago.com.br/developers/en/reference/chargebacks/_payments_id_refunds_refund_id/get
	Get(ctx context.Context, paymentID, refundID int64) (*Response, error)

	// List gets a refund list by payment id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds
	// Reference: https://www.mercadopago.com.br/developers/en/reference/chargebacks/_payments_id_refunds/get
	List(ctx context.Context, paymentID int64) ([]Response, error)

	// Create create a refund by payment id.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds
	// Reference: https://www.mercadopago.com.br/developers/en/reference/chargebacks/_payments_id_refunds/post
	Create(ctx context.Context, request Request, paymentID int64) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	config *config.Config
}

// NewClient returns a new Payment's refund API Client.
func NewClient(c *config.Config) Client {
	return &client{
		config: c,
	}
}

func (c *client) Create(ctx context.Context, request Request, paymentID int64) (*Response, error) {
	conv := strconv.Itoa(int(paymentID))

	body, err := json.Marshal(&request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.Replace(postURL, "{id}", conv, 1), strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var refund *Response
	if err := json.Unmarshal(res, &refund); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return refund, nil
}

func (c *client) Get(ctx context.Context, paymentID, refundID int64) (*Response, error) {
	conv := strconv.Itoa(int(paymentID))
	refundConv := strconv.Itoa(int(refundID))

	url := strings.NewReplacer("{id}", conv, "{refund_id}", refundConv).Replace(getURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var refund *Response
	if err := json.Unmarshal(res, &refund); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return refund, nil
}

func (c *client) List(ctx context.Context, paymentID int64) ([]Response, error) {
	conv := strconv.Itoa(int(paymentID))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.Replace(listURL, "{id}", conv, 1), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var refunds []Response
	if err := json.Unmarshal(res, &refunds); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return refunds, nil
}
