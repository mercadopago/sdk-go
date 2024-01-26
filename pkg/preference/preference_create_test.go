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

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
	"github.com/mercadopago/sdk-go/pkg/option"
)

var (
	createResponseJSON, _ = os.Open("../../resources/mocks/preference/preference_create.json")
	createResponse, _     = io.ReadAll(createResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		credential *credential.Credential
		config     *option.ClientOptions
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
			name: "should_return_error_when_creating_request",
			fields: fields{
				credential: cdt,
				config:     option.ApplyClientOptions(),
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: "error creating request: net/http: nil Context",
		},
		{
			name: "should_return_error_when_send_request",
			fields: fields{
				credential: cdt,
				config: option.ApplyClientOptions(
					option.WithCustomClient(
						&httpclient.Mock{
							DoMock: func(req *http.Request) (*http.Response, error) {
								return nil, fmt.Errorf("some error")
							},
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_error_unmarshal_response",
			fields: fields{
				credential: cdt,
				config: option.ApplyClientOptions(
					option.WithCustomClient(
						&httpclient.Mock{
							DoMock: func(req *http.Request) (*http.Response, error) {
								stringReader := strings.NewReader("invalid json")
								stringReadCloser := io.NopCloser(stringReader)
								return &http.Response{
									Body: stringReadCloser,
								}, nil
							},
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: "invalid character 'i' looking for beginning of value",
		},
		{
			name: "should_return_formatted_response",
			fields: fields{
				credential: cdt,
				config: option.ApplyClientOptions(
					option.WithCustomClient(
						&httpclient.Mock{
							DoMock: func(req *http.Request) (*http.Response, error) {
								stringReader := strings.NewReader(string(createResponse))
								stringReadCloser := io.NopCloser(stringReader)
								return &http.Response{
									Body: stringReadCloser,
								}, nil
							},
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &Response{
					ID:        "1273205088-6da17b39-3bf4-4307-9124-4381c56372e0",
					ClientID: "4679935697572392",
					CollectorID: 1273205088,
					OperationType: "regular_payment",
					DateCreated: parseTime("2024-01-26T08:18:23.229-04:00"),
					Marketplace: "NONE",
					SandboxInitPoint: "https://sandbox.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-6da17b39-3bf4-4307-9124-4381c56372e0",
					SiteID: "MLB",
					BinaryMode: false,
					Expires: false,
					InitPoint: "https://www.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-6da17b39-3bf4-4307-9124-4381c56372e0",
					Items: []PreferenceItem{
						{
							ID:          "123",
							CurrencyID:  "BRL",
							Description: "Description",
							Title:       "Title",
							Quantity:    1,
							UnitPrice:   100,
						},
					},
					PaymentMethods: PreferencePaymentMethods{
						ExcludedPaymentMethods: []PreferencePaymentMethod{
							{},
						},
						ExcludedPaymentTypes: []PreferencePaymentType{
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
				credential: tt.fields.credential,
				config:     tt.fields.config,
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

func parseTime(s string) time.Time {
    t, err := time.Parse(time.RFC3339, s)
    if err != nil {
        panic(err)
    }
    return t
}