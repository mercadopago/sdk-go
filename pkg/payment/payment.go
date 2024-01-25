package payment

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
	postURL   = "https://api.mercadopago.com/v1/payments"
	searchURL = "https://api.mercadopago.com/v1/payments/search"
	getURL    = "https://api.mercadopago.com/v1/payments/{id}"
	putURL    = "https://api.mercadopago.com/v1/payments/{id}"
)

// Client contains the methods to interact with the Payments API.
type Client interface {
	// Create creates a new payment.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/payments
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/payments/_payments/post/
	Create(ctx context.Context, dto Request) (*Response, error)

	// Search searches for payments.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/search
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/payments/_payments_search/get/
	Search(ctx context.Context, f Filters) (*SearchResponse, error)

	// Get gets a payment by its ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/payments/_payments_id/get/
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
	credential *credential.Credential
	config     *option.ClientOptions
}

// NewClient returns a new Payments API Client.
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postURL, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
}

func (c *client) Search(ctx context.Context, f Filters) (*SearchResponse, error) {
	params := url.Values{}
	params.Add("sort", f.Sort)
	params.Add("criteria", f.Criteria)
	params.Add("external_reference", f.ExternalReference)
	params.Add("range", f.Range)
	params.Add("begin_date", f.BeginDate)
	params.Add("end_date", f.EndDate)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
	if err != nil {
		return nil, err
	}

	var formatted *SearchResponse
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
}

func (c *client) Get(ctx context.Context, id int64) (*Response, error) {
	conv := strconv.Itoa(int(id))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.Replace(getURL, "{id}", conv, 1), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
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

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
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

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
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

	res, err := httpclient.Send(ctx, c.credential, c.config.Requester, req)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return formatted, nil
}
