package order

import (
	"context"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
	"net/http"
)

const (
	urlBase           = "https://api.mercadopago.com/v1/orders"
	urlWithOrderID    = urlBase + "/" + "{orderID}"
	urlTransaction    = urlWithOrderID + "/transactions"
	urlProcess        = urlWithOrderID + "/process"
	urlPutTransaction = urlTransaction + "/{transactionID}"
)

// Client contains the methods to interact with the Order API.
type Client interface {
	// Create creates a new order.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/orders
	// Reference: https://www.mercadopago.com.br/developers/pt/reference/order/online-payments/create/post
	Create(ctx context.Context, request Request) (*Response, error)
	Get(ctx context.Context, orderID string) (*Response, error)
	CreateTransaction(ctx context.Context, orderID string, request TransactionRequest) (*TransactionResponse, error)
	Process(ctx context.Context, orderID string) (*Response, error)
	UpdateTransaction(ctx context.Context, orderID string, transactionID string, request PaymentRequest) (*PaymentResponse, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Order API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlBase,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, orderID string) (*Response, error) {
	pathParams := map[string]string{
		"orderID": orderID,
	}
	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlWithOrderID,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Process(ctx context.Context, orderID string) (*Response, error) {
	pathParams := map[string]string{
		"orderID": orderID,
	}
	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlProcess,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) CreateTransaction(ctx context.Context, orderID string, request TransactionRequest) (*TransactionResponse, error) {
	pathParams := map[string]string{
		"orderID": orderID,
	}
	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPost,
		URL:        urlTransaction,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*TransactionResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (c *client) UpdateTransaction(ctx context.Context, orderID string, transactionID string, request PaymentRequest) (*PaymentResponse, error) {
	pathParams := map[string]string{
		"orderID":       orderID,
		"transactionID": transactionID,
	}
	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPut,
		URL:        urlPutTransaction,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*PaymentResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
