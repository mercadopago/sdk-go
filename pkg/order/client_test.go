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
	customerID                                             = "1234567890"
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
				CaptureMode:       "automatic",
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
				Payer: PayerResponse{
					Email:     "{email}",
					FirstName: "John",
					LastName:  "Doe",
					Identification: &IdentificationResponse{
						Type:   "CPF",
						Number: "00000000000",
					},
					Phone: &PhoneResponse{
						AreaCode: "55",
						Number:   "99999999999",
					},
					Address: &AddressResponse{
						StreetName:   "Av. das Nações Unidas",
						StreetNumber: "99",
					},
				},
				Items: []ItemsResponse{
					{
						ID:          "item_id",
						Title:       "Some item title",
						UnitPrice:   "1000.00",
						Description: "Some item description",
						CategoryID:  "category_id",
						Type:        "item_type",
						PictureUrl:  "https://mysite.com/img/item.jpg",
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
				ClientID:          "4599991948843755",
				CollectorID:       "1245621468",
				CreatedDate:       "2024-09-02T22:04:01.880469Z",
				LastUpdatedDate:   "2024-09-02T22:04:04.429289Z",
				Type:              "online",
				Status:            "processed",
				StatusDetail:      "accredited",
				CaptureMode:       "automatic",
				Payer: PayerResponse{
					CustomerID: &customerID,
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
							AutomaticPayments: &AutomaticPaymentResponse{
								PaymentProfileID: "035979dc46c645c4ae12554835b07943",
							},
							StoredCredential: &StoredCredentialResponse{
								PaymentInitiator: "customer",
							},
							SubscriptionData: &SubscriptionDataResponse{
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
				ClientID:          "4599991948843755",
				CollectorID:       "1245621468",
				CreatedDate:       "2024-09-02T22:04:01.880469Z",
				LastUpdatedDate:   "2024-09-02T22:04:04.429289Z",
				Type:              "online",
				Status:            "processing",
				StatusDetail:      "processing",
				CaptureMode:       "automatic",
				Payer: PayerResponse{
					CustomerID: &customerID,
				},
				Transactions: TransactionResponse{
					Payments: []PaymentResponse{
						{
							ID:           "pay_01J6TC8BYRR0T4ZKY0QRTZ0E24",
							ReferenceID:  "74e9f7137a3246d3a3ad607c82da9880",
							Amount:       "200.00",
							Status:       "processing",
							StatusDetail: "processing",
							AutomaticPayments: &AutomaticPaymentResponse{
								PaymentProfileID: "035979dc46c645c4ae12554835b07943",
								Retries:          3,
								ScheduleDate:     "2024-09-02T22:04:01.880469Z",
								DueDate:          "2024-09-02T22:04:01.880469Z",
							},
							StoredCredential: &StoredCredentialResponse{
								PaymentInitiator:   "customer",
								Reason:             "card_on_file",
								StorePaymentMethod: true,
								FirstPayment:       false,
							},
							SubscriptionData: &SubscriptionDataResponse{
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
