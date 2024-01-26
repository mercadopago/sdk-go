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

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
	"github.com/mercadopago/sdk-go/pkg/option"
)

var (
	cdt, _ = credential.New("any")

	getResponseJSON, _ = os.Open("../../resources/mocks/preference/preference_get.json")
	getResponse, _     = io.ReadAll(getResponseJSON)
)

func TestGet(t *testing.T) {
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
								stringReader := strings.NewReader(string(getResponse))
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
					ClientID: "4679935697572392",
					CollectorID: 1273205088,
					OperationType: "regular_payment",
					DateCreated: parseTime("2024-01-23T13:25:20.360-04:00"),
					LastUpdated: parseTime("2024-01-24T17:39:54.750-04:00"),
					Marketplace: "NONE",
					SandboxInitPoint: "https://sandbox.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4",
					SiteID: "MLB",
					BinaryMode: false,
					Expires: false,
					PaymentMethods: PreferencePaymentMethods{
						ExcludedPaymentMethods: []PreferencePaymentMethod{
							{},
						},
						ExcludedPaymentTypes: []PreferencePaymentType{
							{},
						},
					},
					Metadata: map[string]interface{}{},
					ID:        "1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4",
    				InitPoint: "https://www.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4",
					Items: []PreferenceItem{
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
				credential: tt.fields.credential,
				config:     tt.fields.config,
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
