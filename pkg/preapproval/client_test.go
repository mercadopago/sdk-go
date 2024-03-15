package preapproval

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

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

var (
	createResponseJSON, _ = os.Open("../../resources/mocks/preapproval/create_response.json")
	createResponse, _     = io.ReadAll(createResponseJSON)
	getResponseJSON, _    = os.Open("../../resources/mocks/preapproval/get_response.json")
	getResponse, _        = io.ReadAll(getResponseJSON)
	updateResponseJSON, _ = os.Open("../../resources/mocks/preapproval/update_response.json")
	updateResponse, _     = io.ReadAll(updateResponseJSON)
	searchResponseJSON, _ = os.Open("../../resources/mocks/preapproval/search_response.json")
	searchResponse, _     = io.ReadAll(searchResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		config *config.Config
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
		wantErr string
	}{
		{
			name: "should_fail_to_send_request",
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
							stringReaderCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReaderCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: &Response{
				ID:                "2c938084726fca480172750000000000",
				PayerID:           123123123,
				PayerEmail:        "",
				BackURL:           "https://www.mercadopago.com.br",
				CollectorID:       123123123,
				ApplicationID:     123123123,
				Status:            "pending",
				Reason:            "reason",
				ExternalReference: "Ref-123",
				DateCreated:       parseDate("2024-03-06T18:10:01.329-04:00"),
				LastModified:      parseDate("2024-03-06T18:10:01.579-04:00"),
				InitPoint:         "https://www.mercadopago.com.br/subscriptions/checkout?preapproval_id=2c938084726fca480172750000000000",
				AutoRecurring: AutoRecurringResponse{
					Frequency:         1,
					FrequencyType:     "months",
					TransactionAmount: 100,
					CurrencyID:        "BRL",
					FreeTrial:         FreeTrialResponse{},
				},
				Summarized: SummarizedResponse{
					Quotas:                0,
					ChargedQuantity:       0,
					PendingChargeQuantity: 0,
					ChargedAmount:         0,
				},
				NextPaymentDate:    parseDate("2024-03-06T18:10:01.000-04:00"),
				PaymentMethodID:    "",
				FirstInvoiceOffset: "",
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got, err := c.Create(tt.args.ctx, tt.args.request)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
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
				id:  "2c938084726fca480172750000000000",
			},
			want: &Response{
				ID:                "2c938084726fca480172750000000000",
				PayerID:           123123123,
				PayerEmail:        "",
				BackURL:           "https://www.mercadopago.com.br",
				CollectorID:       123123123,
				ApplicationID:     123123123,
				Status:            "pending",
				Reason:            "reason",
				ExternalReference: "Ref-123",
				DateCreated:       parseDate("2024-03-06T18:10:01.329-04:00"),
				LastModified:      parseDate("2024-03-06T18:10:01.579-04:00"),
				InitPoint:         "https://www.mercadopago.com.br/subscriptions/checkout?preapproval_id=2c938084726fca480172750000000000",
				AutoRecurring: AutoRecurringResponse{
					Frequency:         1,
					FrequencyType:     "months",
					TransactionAmount: 100,
					CurrencyID:        "BRL",
					FreeTrial:         FreeTrialResponse{},
				},
				Summarized: SummarizedResponse{
					Quotas:                0,
					ChargedQuantity:       0,
					PendingChargeQuantity: 0,
					ChargedAmount:         0,
				},
				NextPaymentDate:    parseDate("2024-03-06T18:10:01.000-04:00"),
				PaymentMethodID:    "",
				FirstInvoiceOffset: "",
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
					Limit:  10,
					Offset: 10,
					Filters: map[string]string{
						"iD": uuid.NewString(),
					},
				},
			},
			want: &SearchResponse{
				Paging: PagingResponse{
					Offset: 0,
					Limit:  10,
					Total:  10,
				},
				Results: []Response{
					{
						ID:                "2c938084726fca480172750000000000",
						Status:            "pending",
						Reason:            "Yoga classes",
						Summarized:        SummarizedResponse{},
						PayerID:           123123123,
						BackURL:           "https://www.mercadopago.com.br",
						CollectorID:       123123123,
						ApplicationID:     123123123,
						ExternalReference: "Ref-123",
						DateCreated:       parseDate("2023-10-10T10:00:32.895-04:00"),
						LastModified:      parseDate("2023-10-10T10:00:32.896-04:00"),
						InitPoint:         "https://www.mercadopago.com.br/subscriptions/checkout?preapproval_id=2c938084726fca480172750000000000",
						AutoRecurring: AutoRecurringResponse{
							Frequency:         1,
							FrequencyType:     "months",
							TransactionAmount: 10.0,
							CurrencyID:        "BRL",
						},
						PayerFirstName: "Test",
						PayerLastName:  "",
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

func TestUpdate(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx     context.Context
		request UpdateRequest
		id      string
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
							stringReaderCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReaderCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "2c938084726fca480172750000000000",
			},
			want: &Response{
				ID:                "2c938084726fca480172750000000000",
				PayerID:           123123123,
				PayerEmail:        "",
				BackURL:           "https://www.mercadopago.com.br",
				CollectorID:       123123123,
				ApplicationID:     123123123,
				Status:            "pending",
				Reason:            "reason",
				ExternalReference: "Ref-123",
				DateCreated:       parseDate("2024-03-06T18:10:01.329-04:00"),
				LastModified:      parseDate("2024-03-07T17:48:00.821-04:00"),
				InitPoint:         "https://www.mercadopago.com.br/subscriptions/checkout?preapproval_id=2c938084726fca480172750000000000",
				AutoRecurring: AutoRecurringResponse{
					Frequency:         1,
					FrequencyType:     "months",
					TransactionAmount: 50.0,
					CurrencyID:        "BRL",
					FreeTrial:         FreeTrialResponse{},
				},
				Summarized: SummarizedResponse{
					Quotas:                0,
					ChargedQuantity:       0,
					PendingChargeQuantity: 0,
					ChargedAmount:         0,
					Semaphore:             "",
				},
				NextPaymentDate:    parseDate("2024-03-06T18:10:01.000-04:00"),
				PaymentMethodID:    "",
				FirstInvoiceOffset: "",
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got, err := c.Update(tt.args.ctx, tt.args.request, tt.args.id)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}
			if gotErr != tt.wantErr {
				t.Errorf("client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func parseDate(s string) time.Time {
	d, _ := time.Parse(time.RFC3339, s)
	return d
}
