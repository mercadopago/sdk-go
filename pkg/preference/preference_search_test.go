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
	searchResponseJSON, _ = os.Open("../../resources/mocks/preference/preference_search.json")
	searchResponse, _     = io.ReadAll(searchResponseJSON)
)

func TestSearch(t *testing.T) {
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
		want    *SearchResponsePage
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
								stringReader := strings.NewReader(string(searchResponse))
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
			want: &SearchResponsePage{
					Total: 1229,
					NextOffset: 22,
					Elements: []SearchResponse {
						{
							ID:        "1273205088-cf5a445f-9b6e-46a6-9ff9-71e60334979e",
							ClientID: "4679935697572392",
							CollectorID: 1273205088,
							OperationType: "regular_payment",
							DateCreated: parseTime("2023-10-31T10:01:03.555-04:00"),
							ExpirationDateFrom: parseTime("2023-10-31T11:01:03.046-03:00"),
							ExpirationDateTo: parseTime("2023-11-30T11:01:03.046-03:00"),
							Expires: true,
							Marketplace: "NONE",
							SiteID: "MLB",
							Items: []string{
								"Title",
							},
							ExternalReference: "c99c3002-00d4-4c0b-ab92-2470607c788c",
							LiveMode: false,
							PayerEmail: "teste_user_39873928@testuser.com",
							ProcessingModes: []string{
								"aggregator",
							},
							ShippingMode: "custom",
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

			dto := SearchRequest{
				Limit: 22,
				Offset: 100,
			}
			got, err := c.Search(tt.args.ctx, dto)
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
