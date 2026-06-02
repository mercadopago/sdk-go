// Package disbursementrefund provides a client for managing refunds on disbursements
// within MercadoPago advanced (split) payments.
//
// Use this package to list existing refunds, refund all disbursements at once, or
// refund a specific disbursement by amount.
package disbursementrefund

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

// Client defines the interface for interacting with the disbursement refunds API.
type Client interface {
	// ListAll retrieves all refunds for a given advanced payment.
	//
	// GET https://api.mercadopago.com/v1/advanced_payments/{id}/refunds
	ListAll(ctx context.Context, advancedPaymentID int) ([]Response, error)

	// CreateAll refunds all disbursements of an advanced payment at once.
	//
	// POST https://api.mercadopago.com/v1/advanced_payments/{id}/refunds
	CreateAll(ctx context.Context, advancedPaymentID int, request Request) (*Response, error)

	// Create refunds a specific disbursement by amount.
	//
	// POST https://api.mercadopago.com/v1/advanced_payments/{id}/disbursements/{disbursement_id}/refunds
	Create(ctx context.Context, advancedPaymentID int, disbursementID int, request Request) (*Response, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates and returns a new Disbursement Refund API [Client].
func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) ListAll(ctx context.Context, advancedPaymentID int) ([]Response, error) {
	url := fmt.Sprintf("https://api.mercadopago.com/v1/advanced_payments/%d/refunds", advancedPaymentID)
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    url,
	}
	resource, err := httpclient.DoRequest[*[]Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}
	return *resource, nil
}

func (c *client) CreateAll(ctx context.Context, advancedPaymentID int, request Request) (*Response, error) {
	url := fmt.Sprintf("https://api.mercadopago.com/v1/advanced_payments/%d/refunds", advancedPaymentID)
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    url,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (c *client) Create(ctx context.Context, advancedPaymentID int, disbursementID int, request Request) (*Response, error) {
	url := fmt.Sprintf(
		"https://api.mercadopago.com/v1/advanced_payments/%d/disbursements/%d/refunds",
		advancedPaymentID,
		disbursementID,
	)
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    url,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}
	return resource, nil
}
