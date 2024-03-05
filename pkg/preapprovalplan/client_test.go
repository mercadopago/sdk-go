package preapprovalplan

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
	createResponseJSON, _ = os.Open("../../resources/mocks/preapproval_plan/create_response.json")
	createResponse, _     = io.ReadAll(createResponseJSON)
	getResponseJSON, _    = os.Open("../../resources/mocks/preapproval_plan/get_response.json")
	getResponse, _        = io.ReadAll(getResponseJSON)
	updateResponseJSON, _ = os.Open("../../resources/mocks/preapproval_plan/update_response.json")
	updateResponse, _     = io.ReadAll(updateResponseJSON)
	searchResponseJSON, _ = os.Open("../../resources/mocks/preapproval_plan/search_response.json")
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
			name: "should_fail_to_unmarshaling_response",
			fields: fields{
				config: &config.Config{
					Requester: &httpclient.Mock{
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
				ID:            "2c938084726fca480172750000000000",
				CollectorID:   100200300,
				ApplicationID: 1234567812345678,
				Status:        "active",
				DateCreated:   parseDate("2024-02-27T17:37:06.459-04:00"),
				LastModified:  parseDate("2024-02-27T17:37:06.459-04:00"),
				InitPoint:     "https://www.mercadopago.com.br/subscriptions/checkout?preapproval_plan_id=2c938084726fca480172750000000000",
				AutoRecurring: AutoRecurringResponse{
					Frequency:         1,
					FrequencyType:     "days",
					TransactionAmount: 5,
					CurrencyID:        "BRL",
				},
				BackURL: "https://www.yoursite.com",
				PaymentMethodsAllowed: PaymentMethodsAllowedResponse{
					PaymentTypes: []PaymentTypeResponse{
						{
							ID: "credit_card",
						},
					},
					PaymentMethods: []PaymentMethodResponse{
						{
							ID: "bolbradesco",
						},
					},
				},
				Reason: "Yoga classes",
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
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() got = %v, want %v", got, tt.want)
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
				ID:            "2c938084726fca480172750000000000",
				CollectorID:   100200300,
				ApplicationID: 1234567812345678,
				Status:        "active",
				DateCreated:   parseDate("2024-02-27T17:37:06.459-04:00"),
				LastModified:  parseDate("2024-02-27T17:37:06.711-04:00"),
				InitPoint:     "https://www.mercadopago.com.br/subscriptions/checkout?preapproval_plan_id=2c938084726fca480172750000000000",
				AutoRecurring: AutoRecurringResponse{
					Frequency:         1,
					FrequencyType:     "days",
					TransactionAmount: 5,
					CurrencyID:        "BRL",
				},
				BackURL: "https://www.yoursite.com",
				PaymentMethodsAllowed: PaymentMethodsAllowedResponse{
					PaymentTypes: []PaymentTypeResponse{
						{
							ID: "credit_card",
						},
					},
					PaymentMethods: []PaymentMethodResponse{
						{
							ID: "bolbradesco",
						},
					},
				},
				Reason: "Yoga classes",
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

func TestUpdate(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx     context.Context
		id      string
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
				id:  "2c938084726fca480172750000000000",
			},
			want: &Response{
				ID:            "2c938084726fca480172750000000000",
				CollectorID:   100200300,
				ApplicationID: 1234567812345678,
				Status:        "active",
				DateCreated:   parseDate("2024-02-27T17:37:06.459-04:00"),
				LastModified:  parseDate("2024-03-02T19:30:37.530-04:00"),
				InitPoint:     "https://www.mercadopago.com.br/subscriptions/checkout?preapproval_plan_id=2c938084726fca480172750000000000",
				AutoRecurring: AutoRecurringResponse{
					Frequency:         1,
					FrequencyType:     "months",
					TransactionAmount: 10,
					CurrencyID:        "BRL",
					Repetitions:       12,
					FreeTrial: FreeTrialResponse{
						Frequency:          1,
						FrequencyType:      "months",
						FirstInvoiceOffset: 30,
					},
					BillingDay:             10,
					BillingDayProportional: false,
				},
				BackURL: "https://www.yoursite.com",
				PaymentMethodsAllowed: PaymentMethodsAllowedResponse{
					PaymentTypes: []PaymentTypeResponse{
						{
							ID: "credit_card",
						},
					},
					PaymentMethods: []PaymentMethodResponse{
						{
							ID: "bolbradesco",
						},
					},
				},
				Reason: "Yoga classes",
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
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() got = %v, want %v", got, tt.want)
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
					Limit: "10",
				},
			},
			want: &SearchResponse{
				Paging: PagingResponse{
					Offset: 0,
					Total:  10,
					Limit:  10,
				},
				Results: []Response{
					{
						ID:            "2c938084726fca480172750000000000",
						CollectorID:   100200300,
						ApplicationID: 1234567812345678,
						Status:        "active",
						DateCreated:   parseDate("2024-02-27T17:37:06.459-04:00"),
						LastModified:  parseDate("2024-02-27T17:37:06.459-04:00"),
						InitPoint:     "https://www.mercadopago.com.br/subscriptions/checkout?preapproval_plan_id=2c938084726fca480172750000000000",
						AutoRecurring: AutoRecurringResponse{
							Frequency:         1,
							FrequencyType:     "days",
							TransactionAmount: 5,
							CurrencyID:        "BRL",
						},
						BackURL: "https://www.yoursite.com",
						PaymentMethodsAllowed: PaymentMethodsAllowedResponse{
							PaymentTypes: []PaymentTypeResponse{
								{
									ID: "credit_card",
								},
							},
							PaymentMethods: []PaymentMethodResponse{
								{
									ID: "bolbradesco",
								},
							},
						},
						Reason: "Yoga classes",
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

func parseDate(s string) *time.Time {
	d, _ := time.Parse(time.RFC3339, s)
	return &d
}
