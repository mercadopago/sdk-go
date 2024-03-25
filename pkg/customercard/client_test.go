package customercard

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
	responseJSON, _ = os.Open("../../resources/mocks/customer_card/card_response.json")
	response, _     = io.ReadAll(responseJSON)

	listResponseJSON, _ = os.Open("../../resources/mocks/customer_card/list_response.json")
	listResponse, _     = io.ReadAll(listResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx        context.Context
		customerID string
		req        Request
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
				ctx:        context.Background(),
				customerID: "any",
				req:        Request{},
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_card_response",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(response))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				customerID: "1111111111-pDci63MBohR7c",
				req:        Request{Token: "938e19c9848b89fb207105b8a17e97ce"},
			},
			want: &Response{
				ID:              "9999999999",
				CustomerID:      "1111111111-pDci63MBohR7c",
				UserID:          "0000000000",
				FirstSixDigits:  "123456",
				LastFourDigits:  "1234",
				ExpirationMonth: 12,
				ExpirationYear:  2025,
				LiveMode:        true,
				DateCreated:     parseDate("2024-02-07T16:28:38.000-04:00"),
				DateLastUpdated: parseDate("2024-02-07T16:31:06.964-04:00"),
				Issuer: IssuerResponse{
					ID:   24,
					Name: "Mastercard",
				},
				Cardholder: CardholderResponse{
					Name: "APRO",
					Identification: IdentificationResponse{
						Number: "19119119100",
						Type:   "CPF",
					},
				},
				AdditionalInfo: AdditionalInfoResponse{
					RequestPublic:        "true",
					APIClientApplication: "traffic-layer",
					APIClientScope:       "mapi-pci-tl",
				},
				PaymentMethod: PaymentMethodResponse{
					ID:              "master",
					Name:            "Mastercard",
					PaymentTypeID:   "credit_card",
					Thumbnail:       "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
					SecureThumbnail: "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
				},
				SecurityCode: SecurityCodeResponse{
					Length:       3,
					CardLocation: "back",
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{tt.fields.config}
			got, err := c.Create(tt.args.ctx, tt.args.customerID, tt.args.req)
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

func TestUpdate(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx        context.Context
		customerID string
		cardID     string
		req        Request
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
				ctx:        context.Background(),
				customerID: "any",
				req:        Request{},
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_card_response",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(response))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				customerID: "1111111111-pDci63MBohR7c",
				cardID:     "9999999999",
				req:        Request{Token: "938e19c9848b89fb207105b8a17e97ce"},
			},
			want: &Response{
				ID:              "9999999999",
				CustomerID:      "1111111111-pDci63MBohR7c",
				UserID:          "0000000000",
				FirstSixDigits:  "123456",
				LastFourDigits:  "1234",
				ExpirationMonth: 12,
				ExpirationYear:  2025,
				LiveMode:        true,
				DateCreated:     parseDate("2024-02-07T16:28:38.000-04:00"),
				DateLastUpdated: parseDate("2024-02-07T16:31:06.964-04:00"),
				Issuer: IssuerResponse{
					ID:   24,
					Name: "Mastercard",
				},
				Cardholder: CardholderResponse{
					Name: "APRO",
					Identification: IdentificationResponse{
						Number: "19119119100",
						Type:   "CPF",
					},
				},
				AdditionalInfo: AdditionalInfoResponse{
					RequestPublic:        "true",
					APIClientApplication: "traffic-layer",
					APIClientScope:       "mapi-pci-tl",
				},
				PaymentMethod: PaymentMethodResponse{
					ID:              "master",
					Name:            "Mastercard",
					PaymentTypeID:   "credit_card",
					Thumbnail:       "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
					SecureThumbnail: "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
				},
				SecurityCode: SecurityCodeResponse{
					Length:       3,
					CardLocation: "back",
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{tt.fields.config}
			got, err := c.Update(tt.args.ctx, tt.args.customerID, tt.args.cardID, tt.args.req)
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

func TestGet(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx        context.Context
		customerID string
		cardID     string
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
				ctx:        context.Background(),
				customerID: "any",
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_card_response",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(response))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				customerID: "1111111111-pDci63MBohR7c",
				cardID:     "9999999999",
			},
			want: &Response{
				ID:              "9999999999",
				CustomerID:      "1111111111-pDci63MBohR7c",
				UserID:          "0000000000",
				FirstSixDigits:  "123456",
				LastFourDigits:  "1234",
				ExpirationMonth: 12,
				ExpirationYear:  2025,
				LiveMode:        true,
				DateCreated:     parseDate("2024-02-07T16:28:38.000-04:00"),
				DateLastUpdated: parseDate("2024-02-07T16:31:06.964-04:00"),
				Issuer: IssuerResponse{
					ID:   24,
					Name: "Mastercard",
				},
				Cardholder: CardholderResponse{
					Name: "APRO",
					Identification: IdentificationResponse{
						Number: "19119119100",
						Type:   "CPF",
					},
				},
				AdditionalInfo: AdditionalInfoResponse{
					RequestPublic:        "true",
					APIClientApplication: "traffic-layer",
					APIClientScope:       "mapi-pci-tl",
				},
				PaymentMethod: PaymentMethodResponse{
					ID:              "master",
					Name:            "Mastercard",
					PaymentTypeID:   "credit_card",
					Thumbnail:       "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
					SecureThumbnail: "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
				},
				SecurityCode: SecurityCodeResponse{
					Length:       3,
					CardLocation: "back",
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{tt.fields.config}
			got, err := c.Get(tt.args.ctx, tt.args.customerID, tt.args.cardID)
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

func TestDelete(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx        context.Context
		customerID string
		cardID     string
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
				ctx:        context.Background(),
				customerID: "any",
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_card_response",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(response))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				customerID: "1111111111-pDci63MBohR7c",
				cardID:     "9999999999",
			},
			want: &Response{
				ID:              "9999999999",
				CustomerID:      "1111111111-pDci63MBohR7c",
				UserID:          "0000000000",
				FirstSixDigits:  "123456",
				LastFourDigits:  "1234",
				ExpirationMonth: 12,
				ExpirationYear:  2025,
				LiveMode:        true,
				DateCreated:     parseDate("2024-02-07T16:28:38.000-04:00"),
				DateLastUpdated: parseDate("2024-02-07T16:31:06.964-04:00"),
				Issuer: IssuerResponse{
					ID:   24,
					Name: "Mastercard",
				},
				Cardholder: CardholderResponse{
					Name: "APRO",
					Identification: IdentificationResponse{
						Number: "19119119100",
						Type:   "CPF",
					},
				},
				AdditionalInfo: AdditionalInfoResponse{
					RequestPublic:        "true",
					APIClientApplication: "traffic-layer",
					APIClientScope:       "mapi-pci-tl",
				},
				PaymentMethod: PaymentMethodResponse{
					ID:              "master",
					Name:            "Mastercard",
					PaymentTypeID:   "credit_card",
					Thumbnail:       "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
					SecureThumbnail: "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
				},
				SecurityCode: SecurityCodeResponse{
					Length:       3,
					CardLocation: "back",
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{tt.fields.config}
			got, err := c.Delete(tt.args.ctx, tt.args.customerID, tt.args.cardID)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx        context.Context
		customerID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Response
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
				ctx:        context.Background(),
				customerID: "any",
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_card_response",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(listResponse))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				customerID: "1111111111-pDci63MBohR7c",
			},
			want: []Response{
				{
					ID:              "9999999999",
					CustomerID:      "1111111111-pDci63MBohR7c",
					UserID:          "0000000000",
					FirstSixDigits:  "123456",
					LastFourDigits:  "1234",
					ExpirationMonth: 12,
					ExpirationYear:  2025,
					LiveMode:        true,
					DateCreated:     parseDate("2024-02-07T16:28:38.000-04:00"),
					DateLastUpdated: parseDate("2024-02-07T16:31:06.964-04:00"),
					Issuer: IssuerResponse{
						ID:   24,
						Name: "Mastercard",
					},
					Cardholder: CardholderResponse{
						Name: "APRO",
						Identification: IdentificationResponse{
							Number: "19119119100",
							Type:   "CPF",
						},
					},
					AdditionalInfo: AdditionalInfoResponse{
						RequestPublic:        "true",
						APIClientApplication: "traffic-layer",
						APIClientScope:       "mapi-pci-tl",
					},
					PaymentMethod: PaymentMethodResponse{
						ID:              "master",
						Name:            "Mastercard",
						PaymentTypeID:   "credit_card",
						Thumbnail:       "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
						SecureThumbnail: "https://http2.mlstatic.com/storage/logos-api-admin/0daa1670-5c81-11ec-ae75-df2bef173be2-xl@2x.png",
					},
					SecurityCode: SecurityCodeResponse{
						Length:       3,
						CardLocation: "back",
					},
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{tt.fields.config}
			got, err := c.List(tt.args.ctx, tt.args.customerID)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.List() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func parseDate(s string) time.Time {
	d, _ := time.Parse(time.RFC3339, s)
	return d
}
