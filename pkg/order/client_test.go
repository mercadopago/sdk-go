package order

import (
	"context"
	"encoding/json"
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
	searchResponseJSON, _                                  = os.Open("../../resources/mocks/order/search_response.json")
	searchResponse, _                                      = io.ReadAll(searchResponseJSON)
)

func TestSearch(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		request SearchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SearchResponse
		wantErr bool
	}{
		{
			name: "should_fail_to_send_request",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx:     context.Background(),
				request: SearchRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_return_response",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(searchResponse))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				request: SearchRequest{
					Limit:  30,
					Offset: 0,
					Filters: map[string]string{
						"begin_date": "2024-01-01T00:00:00Z",
						"end_date":   "2024-12-31T23:59:59Z",
						"status":     "processed",
					},
				},
			},
			want: &SearchResponse{
				Data: []Response{
					{
						ID:           "01JDNPJ7XJSV1YS2W3JXEJNQ8",
						Status:       "processed",
						StatusDetail: "accredited",
						TotalAmount:  "100.00",
						Currency:     "BRL",
					},
				},
				Paging: &PagingResponse{
					Total:      "1",
					TotalPages: "1",
					Offset:     "0",
					Limit:      "30",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{cfg: tt.fields.cfg}
			got, err := c.Search(tt.args.ctx, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("client.Search() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			name: "should_accept_payment_method_id_bolbradesco",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{"status": "processed"}`
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
				ctx: context.Background(),
				request: Request{
					Type:        "ticket",
					TotalAmount: "84.00",
					Transactions: &TransactionRequest{
						Payments: []PaymentRequest{
							{
								Amount: "84.00",
								PaymentMethod: &PaymentMethodRequest{
									Type: "ticket",
									ID:   "bolbradesco",
								},
							},
						},
					},
				},
			},
			want: &Response{
				Status: "processed",
			},
			wantErr: false,
		},
		{
			name: "should_accept_payment_method_id_boleto",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{"status": "processed"}`
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
				ctx: context.Background(),
				request: Request{
					Type:        "ticket",
					TotalAmount: "42.00",
					Transactions: &TransactionRequest{
						Payments: []PaymentRequest{
							{
								Amount: "42.00",
								PaymentMethod: &PaymentMethodRequest{
									Type: "ticket",
									ID:   "boleto",
								},
							},
						},
					},
				},
			},
			want: &Response{
				Status: "processed",
			},
			wantErr: false,
		},
		{
			name: "should_accept_additional_info_with_multiple_fields",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{"status": "processed"}`
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
				ctx: context.Background(),
				request: Request{
					CaptureMode: "automatic_async",
					Type:        "online",
					TotalAmount: "150.00",
					AdditionalInfo: &AdditionalInfoRequest{
						PayerAuthenticationType: "app",
						PlatformSellerID:        "seller-1234",
						ShipmentExpress:         true,
						PlatformSellerEmail:     "seller@mail.com",
					},
				},
			},
			want: &Response{
				Status: "processed",
			},
			wantErr: false,
		},
		{
			name: "should_accept_capture_mode_automatic_async",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							body := `{"status": "processed"}`
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
				ctx: context.Background(),
				request: Request{
					CaptureMode: "automatic_async",
					Type:        "online",
					TotalAmount: "100.00",
				},
			},
			want: &Response{
				Status: "processed",
			},
			wantErr: false,
		},
		{
			name: "should_accept_capture_mode_manual",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							// Simula um response simples só indicando sucesso
							body := `{"status": "processed"}`
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
				ctx: context.Background(),
				request: Request{
					CaptureMode: "manual",
					Type:        "online",
					TotalAmount: "100.00",
				},
			},
			want: &Response{
				Status: "processed",
			},
			wantErr: false,
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

func TestCheckoutProRequestJSON(t *testing.T) {
	localPickup := true
	freeShipping := true
	maxInstallments := 12
	travelPassengers := []TravelPassengerRequest{
		{
			FirstName:            "John",
			LastName:             "Smith",
			IdentificationType:   "CPF",
			IdentificationNumber: "12345678909",
			ItemReferences:       []string{"ITEM-001"},
		},
	}
	travelRoutes := []TravelRouterRequest{
		{
			Departure:         "SAO",
			Destination:       "RIO",
			DepartureDateTime: "2026-03-10T08:00:00.000-03:00",
			ArrivalDateTime:   "2026-03-10T09:00:00.000-03:00",
			Company:           "TAM",
			ItemReferences:    []string{"ITEM-001"},
		},
	}

	request := Request{
		Type:              "online",
		TotalAmount:       "500.00",
		ExternalReference: "ext_ref_manual_full_test",
		ProcessingMode:    "manual",
		CaptureMode:       "automatic",
		MarketPlaceFee:    "5.00",
		Description:       "Travel package SAO-RIO with insurance",
		ExpirationTime:    "P1D",
		Payer: &PayerRequest{
			Email:     "buyer@mercadopago.com",
			FirstName: "John",
			LastName:  "Smith",
			Phone: &PhoneRequest{
				AreaCode: "11",
				Number:   "999998888",
			},
			Identification: &IdentificationRequest{
				Type:   "CPF",
				Number: "12345678909",
			},
			Address: &PayerAddressRequest{
				ZipCode:      "01310-100",
				StreetName:   "Av. Paulista",
				StreetNumber: "1000",
				Neighborhood: "Bela Vista",
				City:         "Sao Paulo",
			},
		},
		Shipment: &ShipmentRequest{
			Mode:         "custom",
			LocalPickup:  &localPickup,
			Cost:         "15.00",
			FreeShipping: &freeShipping,
			FreeMethods: []FreeMethodRequest{
				{ID: 73328},
			},
			Address: &AddressRequest{
				ZipCode:      "01310-100",
				StreetName:   "Av. Paulista",
				StreetNumber: "1000",
				Floor:        "3",
				Apartment:    "B",
				Neighborhood: "Bela Vista",
				City:         "Sao Paulo",
			},
		},
		Config: &ConfigRequest{
			NotificationURL:       "https://example.com/notifications",
			StatementDescriptor:   "MYSTORE",
			DefaultPaymentDueDate: "P1D",
			Online: &OnlineConfigRequest{
				AvailableFrom:   "2026-01-01T00:00:00Z",
				AllowedUserType: "account_only",
				SuccessURL:      "https://example.com/success",
				FailureURL:      "https://example.com/failure",
				PendingURL:      "https://example.com/pending",
				AutoReturn:      "approved",
				Tracks: []TrackRequest{
					{
						Type: "google_ad",
						Values: map[string]string{
							"conversion_id":    "21312312312123",
							"conversion_label": "TEST",
						},
					},
				},
			},
			PaymentMethod: &PaymentMethodConfigRequest{
				MaxInstallments: &maxInstallments,
				NotAllowedIDs:   []string{"amex"},
				NotAllowedTypes: []string{"ticket"},
				Installments: &InstallmentsConfigRequest{
					InterestFree: &InstallmentsInterestFreeRequest{
						Type:   "range",
						Values: []int{2, 6},
					},
				},
			},
		},
		Items: []ItemsRequest{
			{
				ExternalCode: "ITEM-001",
				Title:        "Flight SAO-RIO",
				Description:  "Round trip, economy class",
				CategoryID:   "travels",
				PictureURL:   "https://example.com/img.jpg",
				Quantity:     1,
				UnitPrice:    "450.00",
				Type:         "travel",
				EventDate:    "2027-01-15T00:00:00.000-03:00",
			},
		},
		AdditionalInfo: &AdditionalInfoRequest{
			PayerRegistrationDate:      "2020-01-15T00:00:00.000-03:00",
			PayerAuthenticationType:    "MOBILE",
			PayerIsPrimeUser:           true,
			PayerIsFirstPurchaseOnLine: true,
			PayerLastPurchase:          "2025-12-01T00:00:00.000-03:00",
			TravelPassengers:           &travelPassengers,
			TravelRoutes:               &travelRoutes,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	config := got["config"].(map[string]any)
	online := config["online"].(map[string]any)
	paymentMethod := config["payment_method"].(map[string]any)
	installments := paymentMethod["installments"].(map[string]any)
	interestFree := installments["interest_free"].(map[string]any)
	shipment := got["shipment"].(map[string]any)
	shipmentAddress := shipment["address"].(map[string]any)
	items := got["items"].([]any)
	firstItem := items[0].(map[string]any)
	additionalInfo := got["additional_info"].(map[string]any)

	if got["type"] != "online" {
		t.Errorf("type = %v, want online", got["type"])
	}
	if got["processing_mode"] != "manual" {
		t.Errorf("processing_mode = %v, want manual", got["processing_mode"])
	}
	if got["total_amount"] != "500.00" {
		t.Errorf("total_amount = %v, want 500.00", got["total_amount"])
	}
	if firstItem["unit_price"] != "450.00" {
		t.Errorf("items[0].unit_price = %v, want 450.00", firstItem["unit_price"])
	}
	if config["notification_url"] != "https://example.com/notifications" {
		t.Errorf("notification_url = %v, want https://example.com/notifications", config["notification_url"])
	}
	if config["statement_descriptor"] != "MYSTORE" {
		t.Errorf("statement_descriptor = %v, want MYSTORE", config["statement_descriptor"])
	}
	if online["available_from"] != "2026-01-01T00:00:00Z" {
		t.Errorf("available_from = %v, want 2026-01-01T00:00:00Z", online["available_from"])
	}
	if online["auto_return"] != "approved" {
		t.Errorf("auto_return = %v, want approved", online["auto_return"])
	}
	if online["success_url"] != "https://example.com/success" {
		t.Errorf("success_url = %v, want https://example.com/success", online["success_url"])
	}
	if online["failure_url"] != "https://example.com/failure" {
		t.Errorf("failure_url = %v, want https://example.com/failure", online["failure_url"])
	}
	if online["pending_url"] != "https://example.com/pending" {
		t.Errorf("pending_url = %v, want https://example.com/pending", online["pending_url"])
	}
	if interestFree["type"] != "range" {
		t.Errorf("interest_free.type = %v, want range", interestFree["type"])
	}
	if shipment["mode"] != "custom" {
		t.Errorf("shipment.mode = %v, want custom", shipment["mode"])
	}
	if shipmentAddress["floor"] != "3" {
		t.Errorf("shipment.address.floor = %v, want 3", shipmentAddress["floor"])
	}
	if shipmentAddress["apartment"] != "B" {
		t.Errorf("shipment.address.apartment = %v, want B", shipmentAddress["apartment"])
	}
	if additionalInfo["payer.authentication_type"] != "MOBILE" {
		t.Errorf("payer.authentication_type = %v, want MOBILE", additionalInfo["payer.authentication_type"])
	}
}

func TestCheckoutProResponseJSON(t *testing.T) {
	body := `{
		"id": "ORDTST01KS5AJ6HTK2HRQ3XJ3C2JCKP9",
		"type": "online",
		"processing_mode": "manual",
		"checkout_url": "https://www.mercadopago.cl/checkout/v1/redirect?order_id=ORDTST01KS5AJ6HTK2HRQ3XJ3C2JCKP9",
		"client_token": "eyJhbGciOiJSUzI1NiIs",
		"config": {
			"statement_descriptor": "MYSTORE",
			"default_payment_due_date": "P1D",
			"online": {
				"available_from": "2026-05-16T18:32:00Z",
				"auto_return": "approved",
				"retries": {
					"allowed": true
				}
			},
			"payment_method": {
				"installments": {
					"interest_free": {
						"type": "range",
						"values": [2, 6]
					}
				}
			}
		}
	}`

	var got Response
	if err := json.Unmarshal([]byte(body), &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if got.CheckoutURL != "https://www.mercadopago.cl/checkout/v1/redirect?order_id=ORDTST01KS5AJ6HTK2HRQ3XJ3C2JCKP9" {
		t.Errorf("CheckoutURL = %v, want checkout redirect URL", got.CheckoutURL)
	}
	if got.ClientToken != "eyJhbGciOiJSUzI1NiIs" {
		t.Errorf("ClientToken = %v, want eyJhbGciOiJSUzI1NiIs", got.ClientToken)
	}
	if got.Config.StatementDescriptor != "MYSTORE" {
		t.Errorf("Config.StatementDescriptor = %v, want MYSTORE", got.Config.StatementDescriptor)
	}
	if got.Config.Online.AvailableFrom != "2026-05-16T18:32:00Z" {
		t.Errorf("Config.Online.AvailableFrom = %v, want 2026-05-16T18:32:00Z", got.Config.Online.AvailableFrom)
	}
	if !got.Config.Online.Retries.Allowed {
		t.Errorf("Config.Online.Retries.Allowed = %v, want true", got.Config.Online.Retries.Allowed)
	}
	if got.Config.PaymentMethodResponse.Installments.InterestFree.Type != "range" {
		t.Errorf("InterestFree.Type = %v, want range", got.Config.PaymentMethodResponse.Installments.InterestFree.Type)
	}
}

func TestOrderRequestJSONIncludesExplicitZeroValues(t *testing.T) {
	localPickup := false
	freeShipping := false
	maxInstallments := 0
	defaultInstallments := 0

	request := Request{
		Shipment: &ShipmentRequest{
			LocalPickup:  &localPickup,
			FreeShipping: &freeShipping,
		},
		Config: &ConfigRequest{
			PaymentMethod: &PaymentMethodConfigRequest{
				MaxInstallments:     &maxInstallments,
				DefaultInstallments: &defaultInstallments,
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	shipment := got["shipment"].(map[string]any)
	config := got["config"].(map[string]any)
	paymentMethod := config["payment_method"].(map[string]any)

	if shipment["local_pickup"] != false {
		t.Errorf("shipment.local_pickup = %v, want false", shipment["local_pickup"])
	}
	if shipment["free_shipping"] != false {
		t.Errorf("shipment.free_shipping = %v, want false", shipment["free_shipping"])
	}
	if paymentMethod["max_installments"] != float64(0) {
		t.Errorf("payment_method.max_installments = %v, want 0", paymentMethod["max_installments"])
	}
	if paymentMethod["default_installments"] != float64(0) {
		t.Errorf("payment_method.default_installments = %v, want 0", paymentMethod["default_installments"])
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

func TestCreate_TransactionSecurity_Simple(t *testing.T) {
	reqBody := Request{
		Config: &ConfigRequest{
			Online: &OnlineConfigRequest{
				TransactionSecurity: &TransactionSecurityRequest{
					Validation:     "always",
					LiabilityShift: "preferred",
				},
			},
		},
	}

	cfg := &config.Config{
		Requester: &httpclient.Mock{
			DoMock: func(req *http.Request) (*http.Response, error) {
				var got struct {
					Config *ConfigRequest `json:"config,omitempty"`
				}
				if err := json.NewDecoder(req.Body).Decode(&got); err != nil {
					return nil, err
				}
				expected := struct {
					Config *ConfigRequest `json:"config,omitempty"`
				}{
					Config: reqBody.Config,
				}
				if !reflect.DeepEqual(got, expected) {
					return nil, fmt.Errorf("body mismatch: got=%+v want=%+v", got, expected)
				}
				return &http.Response{Body: io.NopCloser(strings.NewReader(`{"id":"ord","status":"processing"}`))}, nil
			},
		},
	}

	c := &client{cfg: cfg}
	if _, err := c.Create(context.Background(), reqBody); err != nil {
		t.Fatalf("Create() error: %v", err)
	}
}
