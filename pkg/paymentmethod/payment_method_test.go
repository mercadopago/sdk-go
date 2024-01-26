package paymentmethod

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
	listResponseJSON, _ = os.Open("../../resources/mocks/payment_method/list_response.json")
	listResponse, _     = io.ReadAll(listResponseJSON)
)

func TestList(t *testing.T) {
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
		want    []Response
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
				ctx: context.Background(),
			},
			want: []Response{
				{
					ID:              "debmaster",
					Name:            "Mastercard Débito",
					PaymentTypeID:   "debit_card",
					Status:          "testing",
					SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif",
					Thumbnail:       "https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif",
				},
				{
					ID:              "cabal",
					Name:            "Cabal",
					PaymentTypeID:   "credit_card",
					Status:          "testing",
					SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/cabal.gif",
					Thumbnail:       "https://www.mercadopago.com/org-img/MP3/API/logos/cabal.gif",
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
			got, err := c.List(tt.args.ctx)
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
