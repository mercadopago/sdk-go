package preference

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

var (
	createResponseJSON, _ = os.Open("../../resources/mocks/preference/preference_create.json")
	createResponse, _     = io.ReadAll(createResponseJSON)

	getResponseJSON, _ = os.Open("../../resources/mocks/preference/preference_get.json")
	getResponse, _     = io.ReadAll(getResponseJSON)

	updateResponseJSON, _ = os.Open("../../resources/mocks/preference/preference_update.json")
	updateResponse, _     = io.ReadAll(updateResponseJSON)

	searchResponseJSON, _ = os.Open("../../resources/mocks/preference/preference_search.json")
	searchResponse, _     = io.ReadAll(searchResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr string
	}{
		{
			name: "should_return_error_when_send_request",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_response",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(createResponse))
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
			},
			want: &Response{
				ID:               "1273205088-6da17b39-3bf4-4307-9124-4381c56372e0",
				ClientID:         "4679935697572392",
				CollectorID:      1273205088,
				OperationType:    "regular_payment",
				DateCreated:      parseDate("2024-01-26T08:18:23.229-04:00"),
				Marketplace:      "NONE",
				SandboxInitPoint: "https://sandbox.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-6da17b39-3bf4-4307-9124-4381c56372e0",
				SiteID:           "MLB",
				BinaryMode:       false,
				Expires:          false,
				InitPoint:        "https://www.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-6da17b39-3bf4-4307-9124-4381c56372e0",
				Items: []PreferenceItemResponse{
					{
						ID:          "123",
						CurrencyID:  "BRL",
						Description: "Description",
						Title:       "Title",
						Quantity:    1,
						UnitPrice:   100,
					},
				},
				PaymentMethods: PreferencePaymentMethodsResponse{
					ExcludedPaymentMethods: []PreferencePaymentMethodResponse{
						{},
					},
					ExcludedPaymentTypes: []PreferencePaymentTypeResponse{
						{},
					},
				},
				Metadata: map[string]interface{}{},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}

			dto := Request{
				Items: []PreferenceItemRequest{
					{
						ID:          "123",
						Title:       "Title",
						UnitPrice:   100,
						Quantity:    1,
						Description: "Description",
					},
				},
			}
			got, err := c.Create(tt.args.ctx, dto)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr string
	}{
		{
			name: "should_return_error_when_send_request",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_response",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(getResponse))
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
			},
			want: &Response{
				ClientID:         "4679935697572392",
				CollectorID:      1273205088,
				OperationType:    "regular_payment",
				DateCreated:      parseDate("2024-01-23T13:25:20.360-04:00"),
				LastUpdated:      parseDate("2024-01-24T17:39:54.750-04:00"),
				Marketplace:      "NONE",
				SandboxInitPoint: "https://sandbox.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4",
				SiteID:           "MLB",
				BinaryMode:       false,
				Expires:          false,
				PaymentMethods: PreferencePaymentMethodsResponse{
					ExcludedPaymentMethods: []PreferencePaymentMethodResponse{
						{},
					},
					ExcludedPaymentTypes: []PreferencePaymentTypeResponse{
						{},
					},
				},
				Metadata:  map[string]interface{}{},
				ID:        "1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4",
				InitPoint: "https://www.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4",
				Items: []PreferenceItemResponse{
					{
						ID:          "123",
						CategoryID:  "",
						CurrencyID:  "BRL",
						Description: "updating Description",
						Title:       "updating Title",
						Quantity:    3,
						UnitPrice:   10,
					},
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got, err := c.Get(tt.args.ctx, "1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4")
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr string
	}{
		{
			name: "should_return_error_when_send_request",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_response",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(updateResponse))
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
			},
			want: &Response{
				ClientID:         "4679935697572392",
				CollectorID:      1273205088,
				OperationType:    "regular_payment",
				DateCreated:      parseDate("2024-01-26T09:50:48.888-04:00"),
				LastUpdated:      parseDate("2024-01-26T10:45:08.102-04:00"),
				Marketplace:      "NONE",
				SandboxInitPoint: "https://sandbox.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-6a2d2fa5-edb8-4d06-90c7-74b756a75f38",
				SiteID:           "MLB",
				BinaryMode:       false,
				Expires:          false,
				PaymentMethods: PreferencePaymentMethodsResponse{
					ExcludedPaymentMethods: []PreferencePaymentMethodResponse{
						{},
					},
					ExcludedPaymentTypes: []PreferencePaymentTypeResponse{
						{},
					},
				},
				Metadata:  map[string]interface{}{},
				ID:        "1273205088-6a2d2fa5-edb8-4d06-90c7-74b756a75f38",
				InitPoint: "https://www.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-6a2d2fa5-edb8-4d06-90c7-74b756a75f38",
				Items: []PreferenceItemResponse{
					{
						ID:          "321",
						CategoryID:  "",
						CurrencyID:  "BRL",
						Description: "Updated Description",
						Title:       "Updated Title",
						Quantity:    1,
						UnitPrice:   10,
					},
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}

			dto := Request{
				Items: []PreferenceItemRequest{
					{
						ID:          "321",
						Description: "Updated Description",
						Title:       "Updated Title",
						Quantity:    1,
						UnitPrice:   10,
					},
				},
			}
			got, err := c.Update(tt.args.ctx, dto, "1273205088-6a2d2fa5-edb8-4d06-90c7-74b756a75f38")
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SearchResponsePage
		wantErr string
	}{
		{
			name: "should_return_error_when_send_request",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_response",
			fields: fields{
				config: &config.Config{
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
			},
			want: &SearchResponsePage{
				Total:      1229,
				NextOffset: 22,
				Elements: []SearchResponse{
					{
						ID:                 "1273205088-cf5a445f-9b6e-46a6-9ff9-71e60334979e",
						ClientID:           "4679935697572392",
						CollectorID:        1273205088,
						OperationType:      "regular_payment",
						DateCreated:        parseDate("2023-10-31T10:01:03.555-04:00"),
						ExpirationDateFrom: parseDate("2023-10-31T11:01:03.046-03:00"),
						ExpirationDateTo:   parseDate("2023-11-30T11:01:03.046-03:00"),
						Expires:            true,
						Marketplace:        "NONE",
						SiteID:             "MLB",
						Items: []string{
							"Title",
						},
						ExternalReference: "c99c3002-00d4-4c0b-ab92-2470607c788c",
						LiveMode:          false,
						PayerEmail:        "teste_user_39873928@testuser.com",
						ProcessingModes: []string{
							"aggregator",
						},
						ShippingMode: "custom",
					},
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}

			dto := SearchRequest{
				Filters: map[string]string{
					"SponSOR_ID": "123",
				},
				Limit:  0,
				Offset: 100,
			}
			got, err := c.Search(tt.args.ctx, dto)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Search() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func parseDate(s string) *time.Time {
	d, _ := time.Parse(time.RFC3339, s)
	return &d
}
