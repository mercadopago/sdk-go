package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	baseURL   = "https://api.mercadopago.com/v1/"
	postURL   = baseURL + "payments"
	searchURL = baseURL + "payments/search"
	getURL    = baseURL + "payments/{id}"
	putURL    = baseURL + "payments/{id}"
)

// Client contains the methods to interact with the Payments API.
type Client interface {
	// Create creates a new payment.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/payments
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments/post/
	Create(ctx context.Context, request Request) (*Response, error)

	// Search searches for payments.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/search
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments_search/get/
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Get gets a payment by its ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments_id/get/
	Get(ctx context.Context, id int64) (*Response, error)

	// Cancel cancels a payment by its ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	Cancel(ctx context.Context, id int64) (*Response, error)

	// Capture captures a payment by its ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	Capture(ctx context.Context, id int64) (*Response, error)

	// CaptureAmount captures amount of a payment by its ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	CaptureAmount(ctx context.Context, id int64, amount float64) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	config *config.Config
}

// NewClient returns a new Payments API Client.
func NewClient(c *config.Config) Client {
	return &client{
		config: c,
	}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postURL, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var payment *Response
	if err := json.Unmarshal(res, &payment); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return payment, nil
}

func (c *client) Search(ctx context.Context, dto SearchRequest) (*SearchResponse, error) {
	params := dto.Parameters()

	url, err := url.Parse(searchURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}
	url.RawQuery = params

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var payment *SearchResponse
	if err := json.Unmarshal(res, &payment); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return payment, nil
}

func (c *client) Get(ctx context.Context, id int64) (*Response, error) {
	conv := strconv.Itoa(int(id))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.Replace(getURL, "{id}", conv, 1), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var payment *Response
	if err := json.Unmarshal(res, &payment); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return payment, nil
}

func (c *client) Cancel(ctx context.Context, id int64) (*Response, error) {
	dto := &CancelRequest{Status: "cancelled"}
	body, err := json.Marshal(dto)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	conv := strconv.Itoa(int(id))
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, strings.Replace(putURL, "{id}", conv, 1), strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var payment *Response
	if err := json.Unmarshal(res, &payment); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return payment, nil
}

func (c *client) Capture(ctx context.Context, id int64) (*Response, error) {
	dto := &CaptureRequest{Capture: true}
	body, err := json.Marshal(dto)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	conv := strconv.Itoa(int(id))
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, strings.Replace(putURL, "{id}", conv, 1), strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var payment *Response
	if err := json.Unmarshal(res, &payment); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return payment, nil
}

func (c *client) CaptureAmount(ctx context.Context, id int64, amount float64) (*Response, error) {
	dto := &CaptureRequest{TransactionAmount: amount, Capture: true}
	body, err := json.Marshal(dto)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	conv := strconv.Itoa(int(id))
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, strings.Replace(putURL, "{id}", conv, 1), strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var payment *Response
	if err := json.Unmarshal(res, &payment); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return payment, nil
}

func buildUrl(params url.Values) (string, error) {
	url, err := url.Parse(searchURL)
	if err != nil {
		return "", fmt.Errorf("error parsing url: %w", err)
	}

	for key, value := range params {
		for _, v := range value {
			q := url.Query()
			q.Add(key, v)
			url.RawQuery = q.Encode()
		}
	}
	return url.String(), nil
}
