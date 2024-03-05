package invoice

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	getResponseJSON, _    = os.Open("../../resources/mocks/invoice/get_response.json")
	getResponse, _        = io.ReadAll(getResponseJSON)
	searchResponseJSON, _ = os.Open("../../resources/mocks/invoice/search_response.json")
	searchResponse, _     = io.ReadAll(searchResponseJSON)
)

func TestGet(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx context.Context
		id  string
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
				id:  "3950169598",
			},
			want: &Response{
				PreapprovalID:     "202caa5d4084417b8e2a394121bf172b",
				ID:                3950169598,
				Type:              "recurring",
				Status:            "processed",
				DateCreated:       parseDate("2024-02-27T17:42:04.835-04:00"),
				LastModified:      parseDate("2024-02-27T17:45:06.462-04:00"),
				TransactionAmount: 5.00,
				CurrencyID:        "BRL",
				Reason:            "Yoga classes",
				Payment: PaymentResponse{
					ID:           3950169598,
					Status:       "approved",
					StatusDetail: "accredited",
				},
				RetryAttempt:    1,
				NextRetryDate:   parseDate("2024-02-28T17:40:33.000-04:00"),
				DebitDate:       parseDate("2024-02-27T17:40:32.000-04:00"),
				PaymentMethodID: "account_money",
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got, err := c.Get(tt.args.ctx, tt.args.id)
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

func TestSearch(t *testing.T) {
	type fields struct {
		config *config.Config
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
				request: SearchRequest{
					Limit: "30",
				},
			},
			want: &SearchResponse{
				Results: []Response{
					{
						PreapprovalID:     "202caa5d4084417b8e2a394121bf172b",
						ID:                3950169598,
						Type:              "recurring",
						Status:            "processed",
						DateCreated:       parseDate("2024-02-27T17:42:04.835-04:00"),
						LastModified:      parseDate("2024-02-27T17:45:06.462-04:00"),
						TransactionAmount: 5.00,
						CurrencyID:        "BRL",
						Reason:            "Yoga classes",
						Payment: PaymentResponse{
							ID:           3950169598,
							Status:       "approved",
							StatusDetail: "accredited",
						},
						RetryAttempt:    1,
						NextRetryDate:   parseDate("2024-02-28T17:40:33.000-04:00"),
						DebitDate:       parseDate("2024-02-27T17:40:32.000-04:00"),
						PaymentMethodID: "account_money",
					},
				},
				Paging: Paging{
					Offset: 0,
					Limit:  12,
					Total:  1,
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
			got, err := c.Search(tt.args.ctx, tt.args.request)
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
