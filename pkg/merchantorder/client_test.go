package merchantorder

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
	getResponseJSON, _ = os.Open("../../resources/mocks/merchantorder/get_response.json")
	getResponse, _     = io.ReadAll(getResponseJSON)

	searchResponseJSON, _ = os.Open("../../resources/mocks/merchantorder/search_response.json")
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
			want:    buildResponseMock(),
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got, err := c.Create(tt.args.ctx, Request{})
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
			want:    buildResponseMock(),
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got, err := c.Update(tt.args.ctx, UpdateRequest{}, 1622029222)
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
			want:    buildResponseMock(),
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got, err := c.Get(tt.args.ctx, 12355555)
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
					Limit:   1,
					Offset:  1,
					Filters: map[string]string{"preference": "134445566"},
				},
			},
			want:    buildSearchResponseMock(),
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

func buildResponseMock() *Response {
	dateCreated, _ := time.Parse(time.RFC3339, "2023-03-28T13:50:19.143-04:00")
	lastUpdate, _ := time.Parse(time.RFC3339, "2023-03-28T13:50:19.143-04:00")

	return &Response{
		ID:           1234545454543,
		Status:       "opened",
		PreferenceID: "1340175910-f2694bdc-7562-499e-a373-057cce3a027b",
		Collector: CollectorResponse{
			ID:       1340175910,
			Nickname: "TEST_USER_658045679",
		},
		Marketplace: "NONE",
		DateCreated: &dateCreated,
		LastUpdated: &lastUpdate,
		TotalAmount: 10,
		SiteID:      "MLB",
		Items: []ItemResponse{
			{
				CurrencyID:  "BRL",
				Description: "Some book",
				Title:       "Produto sem descrição",
				Quantity:    1,
				UnitPrice:   10,
			},
		},
		OrderStatus: "payment_required",
	}
}

func buildSearchResponseMock() *SearchResponse {
	return &SearchResponse{
		Elements: []Response{*buildResponseMock()},
		Total:    1,
		Offset:   1,
	}
}
