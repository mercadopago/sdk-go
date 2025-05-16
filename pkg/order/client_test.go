package order

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

var (
	createResponseJSON, _                                  = os.Open("../../resources/mocks/order/create_full_order_response.json")
	createResponse, _                                      = io.ReadAll(createResponseJSON)
	createAutomaticPaymentsSyncResponseJSON, _             = os.Open("../../resources/mocks/order/create_order_automatic_payments_sync_response.json")
	createAutomaticPaymentsSyncResponse, _                 = io.ReadAll(createAutomaticPaymentsSyncResponseJSON)
	createAutomaticPaymentsWithProfileAsyncResponseJSON, _ = os.Open("../../resources/mocks/order/create_order_automatic_payments_with_profile_async_response.json")
	createAutomaticPaymentsWithProfileAsyncResponse, _     = io.ReadAll(createAutomaticPaymentsWithProfileAsyncResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		request Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "should_fail_to_send_request",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error")
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_return_full_order_response",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(createResponse))
							stringReaderCloser := io.NopCloser(stringReader)
							return &http.Response{Body: stringReaderCloser}, nil
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: &Response{
				ID:                "01HRYFWNYRE1MR1E60MW3X0T2P",
				Type:              "online",
				TotalAmount:       "1000.00",
				ExternalReference: "ext_ref_1234",
				CountryCode:       "ARG",
				Status:            "processed",
				StatusDetail:      "accredited",
				CaptureMode:       "automatic_async",
				ProcessingMode:    "automatic",
				Description:       "some description",
				Marketplace:       "NONE",
				MarketplaceFee:    "10.00",
				ExpirationTime:    "P3D",
				Transactions: TransactionResponse{
					Payments: []PaymentResponse{
						{
							ID:          "01HRYFXQ53Q3JPEC48MYWMR0TE",
							ReferenceID: "123456789",
							Status:      "processed",
							Amount:      "1000.00",
							PaymentMethod: PaymentMethodResponse{
								ID:                  "master",
								Type:                "credit_card",
								Token:               "677859ef5f18ea7e3a87c41d02c3fbe3",
								StatementDescriptor: "LOJA X",
								Installments:        1,
							},
						},
					},
					Refunds: []RefundResponse{},
				},
				Items: []ItemsResponse{
					{
						Title:       "Some item title",
						UnitPrice:   "1000.00",
						Description: "Some item description",
						CategoryID:  "category_id",
						Type:        "item_type",
						PictureURL:  "https://mysite.com/img/item.jpg",
						Quantity:    1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "should_return_order_automatic_payments_sync_response",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(createAutomaticPaymentsSyncResponse))
							stringReaderCloser := io.NopCloser(stringReader)
							return &http.Response{Body: stringReaderCloser}, nil
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: &Response{
				ID:                "01J6TC8BYRR0T4ZKY0QR39WGYE",
				ProcessingMode:    "automatic",
				ExternalReference: "ext_ref_1234",
				Marketplace:       "NONE",
				TotalAmount:       "200.00",
				CountryCode:       "MEX",
				Type:              "online",
				Status:            "processed",
				StatusDetail:      "accredited",
				CaptureMode:       "automatic_async",
				CreatedDate:       "2024-09-02T22:04:01.880469Z",
				LastUpdatedDate:   "2024-09-02T22:04:04.429289Z",
				Payer: PayerResponse{
					CustomerID: "1234567890",
				},
				Transactions: TransactionResponse{
					Payments: []PaymentResponse{
						{
							ID:           "pay_01J6TC8BYRR0T4ZKY0QRTZ0E24",
							ReferenceID:  "74e9f7137a3246d3a3ad607c82da9880",
							Amount:       "200.00",
							Status:       "processed",
							StatusDetail: "accredited",
							PaymentMethod: PaymentMethodResponse{
								ID:           "master",
								CardID:       "9514636140",
								Type:         "credit_card",
								Installments: 1,
								Token:        "677859ef5f18ea7e3a87c41d02c3fbe3",
							},
							AutomaticPayments: AutomaticPaymentResponse{
								PaymentProfileID: "035979dc46c645c4ae12554835b07943",
							},
							StoredCredential: StoredCredentialResponse{
								PaymentInitiator: "customer",
							},
							SubscriptionData: SubscriptionDataResponse{
								SubscriptionSequence: SubscriptionSequenceResponse{
									Number: 1,
									Total:  12,
								},
								InvoiceID: "00000000000",
								InvoicePeriod: InvoicePeriodResponse{
									Period: 1,
									Type:   "month",
								},
								BillingDate: "2024-09-02",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "should_return_order_automatic_payments_with_profile_async_response",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(createAutomaticPaymentsWithProfileAsyncResponse))
							stringReaderCloser := io.NopCloser(stringReader)
							return &http.Response{Body: stringReaderCloser}, nil
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: &Response{
				ID:                "01J6TC8BYRR0T4ZKY0QR39WGYE",
				ProcessingMode:    "automatic_async",
				ExternalReference: "ext_ref_1234",
				Marketplace:       "NONE",
				TotalAmount:       "200.00",
				CountryCode:       "MEX",
				Type:              "online",
				Status:            "processing",
				StatusDetail:      "processing",
				CaptureMode:       "automatic_async",
				CreatedDate:       "2024-09-02T22:04:01.880469Z",
				LastUpdatedDate:   "2024-09-02T22:04:04.429289Z",
				Payer: PayerResponse{
					CustomerID: "1234567890",
				},
				Transactions: TransactionResponse{
					Payments: []PaymentResponse{
						{
							ID:           "pay_01J6TC8BYRR0T4ZKY0QRTZ0E24",
							ReferenceID:  "74e9f7137a3246d3a3ad607c82da9880",
							Amount:       "200.00",
							Status:       "processing",
							StatusDetail: "processing",
							AutomaticPayments: AutomaticPaymentResponse{
								PaymentProfileID: "035979dc46c645c4ae12554835b07943",
								Retries:          3,
								ScheduleDate:     "2024-09-02T22:04:01.880469Z",
								DueDate:          "2024-09-02T22:04:01.880469Z",
							},
							StoredCredential: StoredCredentialResponse{
								PaymentInitiator:   "customer",
								Reason:             "card_on_file",
								StorePaymentMethod: true,
								FirstPayment:       false,
							},
							SubscriptionData: SubscriptionDataResponse{
								SubscriptionSequence: SubscriptionSequenceResponse{
									Number: 1,
									Total:  12,
								},
								InvoiceID: "00000000000",
								InvoicePeriod: InvoicePeriodResponse{
									Period: 1,
									Type:   "month",
								},
								BillingDate: "2024-09-02",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Create(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrder(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		orderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "should_fail_to_get_request",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error")
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "invalidOrderID",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_succeed_to_get_request",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{"id": "validOrderID", "status": "processed"}`
							resp := &http.Response{
								StatusCode: http.StatusOK,
								Body:       io.NopCloser(strings.NewReader(body)),
								Header:     make(http.Header),
							}
							return resp, nil
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "validOrderID",
			},
			want: &Response{
				ID:     "validOrderID",
				Status: "processed",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Get(tt.args.ctx, tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestProcessOrder(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		orderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "should_fail_to_process_order",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error processing order")
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "invalidOrderID",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_succeed_to_process_order",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{
						"id": "ORD01JKV9XT88XQJ7H5W83C2MA5EP",
						"status": "processed",
						"status_detail": "accredited"
					}` // Removi a vírgula extra aqui
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       io.NopCloser(strings.NewReader(body)),
								Header:     make(http.Header),
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "validOrderID",
			},
			want: &Response{
				ID:           "ORD01JKV9XT88XQJ7H5W83C2MA5EP",
				Status:       "processed",
				StatusDetail: "accredited",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Process(tt.args.ctx, tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestCreateTransaction(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		orderID string
		request TransactionRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TransactionResponse
		wantErr bool
	}{
		{
			name: "should_fail_to_create_transaction",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error creating transaction")
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "Order123",
				request: TransactionRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_succeed_to_create_transaction",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{
                                "payments": [
                                    {
                                        "id": "payment_12345",
                                        "amount": "100.00",
                                        "currency": "BRL",
                                        "payment_method": {
                                            "id": "master",
                                            "type": "credit_card",
                                            "token": "token_1234",
                                            "installments": 1,
                                            "statement_descriptor": "statement"
                                        }
                                    }
                                ]
                            }`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       io.NopCloser(strings.NewReader(body)),
								Header:     make(http.Header),
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "validOrderID",
				request: TransactionRequest{
					Payments: []PaymentRequest{
						{
							Amount: "100.00",
							PaymentMethod: &PaymentMethodRequest{
								ID:                  "master",
								Type:                "credit_card",
								Token:               "token_1234",
								Installments:        1,
								StatementDescriptor: "statement",
							},
						},
					},
				},
			},
			want: &TransactionResponse{
				Payments: []PaymentResponse{
					{
						ID:     "payment_12345",
						Amount: "100.00",
						PaymentMethod: PaymentMethodResponse{
							ID:                  "master",
							Type:                "credit_card",
							Token:               "token_1234",
							Installments:        1,
							StatementDescriptor: "statement",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.CreateTransaction(tt.args.ctx, tt.args.orderID, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create Transaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create Transaction() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestUpdateTransaction(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx           context.Context
		orderID       string
		transactionID string
		request       PaymentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PaymentResponse
		wantErr bool
	}{
		{
			name: "should_fail_to_update_transaction",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error updating transaction")
						},
					},
				},
			},
			args: args{
				ctx:           context.Background(),
				orderID:       "Order123",
				transactionID: "Pay_Order123",
				request:       PaymentRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_succeed_to_update_transaction",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{
										"payment_method": {
											"type": "credit_card",
											"installments": 2,
											"statement_descriptor": "updated statement"
										}
									}`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       io.NopCloser(strings.NewReader(body)),
								Header:     make(http.Header),
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:           context.Background(),
				orderID:       "Order123",
				transactionID: "Pay_Order123",
				request: PaymentRequest{
					PaymentMethod: &PaymentMethodRequest{
						Type:                "credit_card",
						Installments:        2,
						StatementDescriptor: "updated statement",
					},
				},
			},
			want: &PaymentResponse{
				PaymentMethod: PaymentMethodResponse{
					Type:                "credit_card",
					Installments:        2,
					StatementDescriptor: "updated statement",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.UpdateTransaction(tt.args.ctx, tt.args.orderID, tt.args.transactionID, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update Transaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update Transaction() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestCancelOrder(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		orderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "should_fail_to_cancel_order",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error canceling order")
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "invalidOrderID",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_succeed_to_cancel_order",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{"id": "validOrderID", "status": "cancelled"}`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       io.NopCloser(strings.NewReader(body)),
								Header:     make(http.Header),
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "validOrderID",
			},
			want: &Response{
				ID:     "validOrderID",
				Status: "cancelled",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Cancel(tt.args.ctx, tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cancel() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestCapture(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		orderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "should_fail_to_capture_order",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error capturing order")
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "invalidOrderID",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_succeed_to_capture_order",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{"id": "validOrderID", "status": "captured"}`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       io.NopCloser(strings.NewReader(body)),
								Header:     make(http.Header),
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "validOrderID",
			},
			want: &Response{
				ID:     "validOrderID",
				Status: "captured",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Capture(tt.args.ctx, tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Capture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Capture() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestRefund(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		orderID string
		request *RefundRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "should_fail_to_refund_order",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error refunding order")
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "invalidOrderID",
				request: &RefundRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_succeed_to_refund_order",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{"id": "validOrderID", "status": "refunded"}`
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       io.NopCloser(strings.NewReader(body)),
								Header:     make(http.Header),
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				orderID: "validOrderID",
				request: nil,
			},
			want: &Response{
				ID:     "validOrderID",
				Status: "refunded",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Refund(tt.args.ctx, tt.args.orderID, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Refund() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Refund() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDeleteTransaction(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx           context.Context
		orderID       string
		transactionID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should_fail_to_delete_transaction",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error deleting transaction")
						},
					},
				},
			},
			args: args{
				ctx:           context.Background(),
				orderID:       "Order123",
				transactionID: "invalidTransactionID",
			},
			wantErr: true,
		},
		{
			name: "should_succeed_to_delete_transaction",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header:     make(http.Header),
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:           context.Background(),
				orderID:       "validOrderID",
				transactionID: "validTransactionID",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			err := c.DeleteTransaction(tt.args.ctx, tt.args.orderID, tt.args.transactionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		{
			name: "should_create_new_client_with_valid_config",
			args: args{
				cfg: &config.Config{
					AccessToken: "test_access_token",
				},
			},
			want: &client{
				cfg: &config.Config{
					AccessToken: "test_access_token",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args.cfg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
