package preference

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"reflect"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

var (
	updateResponseJSON, _ = os.Open("../../resources/mocks/preference/preference_update.json")
	updateResponse, _     = io.ReadAll(updateResponseJSON)
)

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
			name: "should_return_error_when_creating_request",
			fields: fields{
				config: nil,
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
				config: &config.Config{
					HTTPClient: &httpclient.Mock{
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
			name: "should_return_error_unmarshal_response",
			fields: fields{
				config: &config.Config{
					HTTPClient: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader("invalid json")
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
			want:    nil,
			wantErr: "invalid character 'i' looking for beginning of value",
		},
		{
			name: "should_return_formatted_response",
			fields: fields{
				config: &config.Config{
					HTTPClient: &httpclient.Mock{
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
				DateCreated:      parseTime("2024-01-26T09:50:48.888-04:00"),
				LastUpdated:      parseTime("2024-01-26T10:45:08.102-04:00"),
				Marketplace:      "NONE",
				SandboxInitPoint: "https://sandbox.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-6a2d2fa5-edb8-4d06-90c7-74b756a75f38",
				SiteID:           "MLB",
				BinaryMode:       false,
				Expires:          false,
				PaymentMethods: PreferencePaymentMethods{
					ExcludedPaymentMethods: []PreferencePaymentMethod{
						{},
					},
					ExcludedPaymentTypes: []PreferencePaymentType{
						{},
					},
				},
				Metadata:  map[string]interface{}{},
				ID:        "1273205088-6a2d2fa5-edb8-4d06-90c7-74b756a75f38",
				InitPoint: "https://www.mercadopago.com.br/checkout/v1/redirect?pref_id=1273205088-6a2d2fa5-edb8-4d06-90c7-74b756a75f38",
				Items: []PreferenceItem{
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
				config: tt.fields.config,
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
			got, err := c.Update(tt.args.ctx, "1273205088-6a2d2fa5-edb8-4d06-90c7-74b756a75f38", dto)
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
