package payment

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
	createResponseJSON, _ = os.Open("../../resources/mocks/payment/create_response.json")
	createResponse, _     = io.ReadAll(createResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx context.Context
		dto Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr string
	}{
		{
			name: "should_fail_to_marshal_dto",
			fields: fields{
				config: nil,
			},
			args: args{
				ctx: nil,
				dto: Request{
					Metadata: map[string]any{
						"fail": make(chan int),
					},
				},
			},
			want:    nil,
			wantErr: "error marshaling request body: json: unsupported type: chan int",
		},
		{
			name: "should_fail_to_create_request",
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
			name: "should_fail_to_send_request",
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
			name: "should_fail_to_unmarshaling_response",
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
			wantErr: "error unmarshaling response: invalid character 'i' looking for beginning of value",
		},
		{
			name: "should_return_formatted_response",
			fields: fields{
				config: &config.Config{
					HTTPClient: &httpclient.Mock{
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
				ID:                 1316782634,
				DateCreated:        func() *time.Time { t := time.Date(2024, 1, 26, 10, 1, 4, 655, time.UTC); return &t }(),
				DateLastUpdated:    func() *time.Time { t := time.Date(2024, 1, 26, 10, 1, 4, 655, time.UTC); return &t }(),
				DateOfExpiration:   nil,
				MoneyReleaseDate:   nil,
				MoneyReleaseStatus: "released",
				OperationType:      "regular_payment",
				IssuerID:           "162",
				PaymentMethodID:    "master",
				PaymentTypeID:      "credit_card",
				PaymentMethod: &PaymentMethodResponse{
					ID:       "master",
					Type:     "credit_card",
					IssuerID: "162",
				},
				// Status: "pending",
				// Status_detail: "pending_challenge",
				// CurrencyID: "MXN",
				// Description: nil,
				// LiveMode: false,
				// SponsorID: nil,
				// authorization_code nil,
				// money_release_schema nil,
				// taxes_amount 0,
				// counter_currency nil,
				// brand_id nil,
				// shipping_amount 0,
				// build_version "54.20.12-hotfix-9",
				// pos_id nil,
				// store_id nil,
				// integrator_id nil,
				// platform_id nil,
				// corporation_id nil,
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				config: tt.fields.config,
			}
			got, err := c.Create(tt.args.ctx, tt.args.dto)
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
