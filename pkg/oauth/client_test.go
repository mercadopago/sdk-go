package oauth

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
	oauthResponseJSON, _ = os.Open("../../resources/mocks/oauth/response.json")
	oauthResponse, _     = io.ReadAll(oauthResponseJSON)
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
							stringReader := strings.NewReader(string(oauthResponse))
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
				AccessToken:  "APP_USR-1223334455",
				Scope:        "offline_access payments read write",
				RefreshToken: "TG-65cf4eed634",
				PublicKey:    "APP_USR-5b5b91b7",
				TokenType:    "Bearer",
				LiveMode:     true,
				ExpiresIn:    15552000,
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got, err := c.Create(tt.args.ctx, "TG-65cf4eed634", "http://test.com")
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

func TestRefresh(t *testing.T) {
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
							stringReader := strings.NewReader(string(oauthResponse))
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
				AccessToken:  "APP_USR-1223334455",
				Scope:        "offline_access payments read write",
				RefreshToken: "TG-65cf4eed634",
				PublicKey:    "APP_USR-5b5b91b7",
				TokenType:    "Bearer",
				LiveMode:     true,
				ExpiresIn:    15552000,
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got, err := c.Refresh(tt.args.ctx, "TG-65cf4eed634")
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Refresh() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Refresh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAuthorizationURL(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		clientID    string
		redirectURI string
		state       string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should_return_authorization_url",
			fields: fields{
				config: &config.Config{
					AccessToken: "accessToken",
				},
			},
			args: args{
				clientID:    "323123123",
				redirectURI: "redirectURI",
				state:       "state",
			},
			want: "https://auth.mercadopago.com/authorization?client_id=323123123&platform_id=mp&redirect_uri=redirectURI&response_type=code&state=state",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.config,
			}
			got := c.GetAuthorizationURL(tt.args.clientID, tt.args.redirectURI, tt.args.state)

			if got != tt.want {
				t.Errorf("client.getAuthorizationURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
