package payment

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
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
	Create(dto Request, opts ...httpclient.RequestOption) (*Response, error)

	// Search searches for payments.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/search
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/payments/_payments_search/get/
	Search(f Filters, opts ...httpclient.RequestOption) (*SearchResponse, error)

	// Get gets a payment by its ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/payments/_payments_id/get/
	Get(id int64, opts ...httpclient.RequestOption) (*Response, error)

	// Cancel cancels a payment by its ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	Cancel(id int64, opts ...httpclient.RequestOption) (*Response, error)

	// Capture captures a payment by its ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	Capture(id int64, opts ...httpclient.RequestOption) (*Response, error)

	// CaptureAmount captures amount of a payment by its ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	CaptureAmount(id int64, amount float64, opts ...httpclient.RequestOption) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	requester httpclient.Requester
}

// NewClient returns a new Payments API Client.
func NewClient(opts ...httpclient.APIOption) Client {
	built := httpclient.BuildAPIClientOptions(opts...)

	return &client{
		requester: built.APIRequester,
	}
}

func (c *client) Create(dto Request, opts ...httpclient.RequestOption) (*Response, error) {
	body, err := json.Marshal(&dto)
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error marshaling request body: " + err.Error(),
		}
	}

	req, err := http.NewRequest(http.MethodPost, postURL, strings.NewReader(string(body)))
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error creating request: " + err.Error(),
		}
	}

	res, err := httpclient.Send(c.requester, req, opts...)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}

func (c *client) Search(f Filters, opts ...httpclient.RequestOption) (*SearchResponse, error) {
	params := url.Values{}
	params.Add("sort", f.Sort)
	params.Add("criteria", f.Criteria)
	params.Add("external_reference", f.ExternalReference)
	params.Add("range", f.Range)
	params.Add("begin_date", f.BeginDate)
	params.Add("end_date", f.EndDate)

	req, err := http.NewRequest(http.MethodGet, searchURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error creating request: " + err.Error(),
		}
	}

	res, err := httpclient.Send(c.requester, req, opts...)
	if err != nil {
		return nil, err
	}

	var formatted *SearchResponse
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}

func (c *client) Get(id int64, opts ...httpclient.RequestOption) (*Response, error) {
	conv := strconv.Itoa(int(id))

	req, err := http.NewRequest(http.MethodGet, strings.Replace(getURL, "{id}", conv, 1), nil)
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error creating request: " + err.Error(),
		}
	}

	res, err := httpclient.Send(c.requester, req, opts...)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}

func (c *client) Cancel(id int64, opts ...httpclient.RequestOption) (*Response, error) {
	dto := &CancelRequest{Status: "cancelled"}
	body, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	conv := strconv.Itoa(int(id))
	req, err := http.NewRequest(http.MethodPut, strings.Replace(putURL, "{id}", conv, 1), strings.NewReader(string(body)))
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error creating request: " + err.Error(),
		}
	}

	res, err := httpclient.Send(c.requester, req, opts...)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}

func (c *client) Capture(id int64, opts ...httpclient.RequestOption) (*Response, error) {
	dto := &CaptureRequest{Capture: true}
	body, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	conv := strconv.Itoa(int(id))
	req, err := http.NewRequest(http.MethodPut, strings.Replace(putURL, "{id}", conv, 1), strings.NewReader(string(body)))
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error creating request: " + err.Error(),
		}
	}

	res, err := httpclient.Send(c.requester, req, opts...)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}

func (c *client) CaptureAmount(id int64, amount float64, opts ...httpclient.RequestOption) (*Response, error) {
	dto := &CaptureRequest{TransactionAmount: amount, Capture: true}
	body, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	conv := strconv.Itoa(int(id))
	req, err := http.NewRequest(http.MethodPut, strings.Replace(putURL, "{id}", conv, 1), strings.NewReader(string(body)))
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error creating request: " + err.Error(),
		}
	}

	res, err := httpclient.Send(c.requester, req, opts...)
	if err != nil {
		return nil, err
	}

	formatted := &Response{}
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}
