package defaultrequester

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	server, req := NewRequestWithHTTPServerUnavailableMock()
	defer server.Close()

	serverOK, reqOK := NewRequestWithHTTPServerOKMock()
	defer serverOK.Close()

	reqWithDeadline, cancel := NewRequestMockWithDeadlineContextAndServerError()
	defer cancel()

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name       string
		args       args
		wantStatus string
		wantErr    string
	}{
		{
			name: "should_return_response_ok_when_status_code_is_200",
			args: args{
				req: reqOK,
			},
			wantStatus: "200 OK",
		},
		{
			name: "should_retry_and_return_response_error_when_status_code_is_503",
			args: args{
				req: req,
			},
			wantStatus: "503 Service Unavailable",
		},
		{
			name: "should_return_error_when_context_is_canceled",
			args: args{
				req: NewRequestMockWithCanceledContext(),
			},
			wantErr: "context canceled",
		},
		{
			name: "should_return_error_when_context_has_deadline_smaller_than_backoff",
			args: args{
				req: reqWithDeadline,
			},
			wantErr: "Get \"\": unsupported protocol scheme \"\"",
		},
		{
			name: "should_return_error_when_retry_is_enabled_and_request_fails",
			args: args{
				req: NewRequestMock(),
			},
			wantErr: "Get \"\": unsupported protocol scheme \"\"",
		},
		{
			name: "should_return_error_when_request_is_nil",
			args: args{
				req: NewInvalidRequestMock(),
			},
			wantErr: "error getting body",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New()
			got, err := d.Do(tt.args.req)

			gotError := ""
			if err != nil {
				gotError = err.Error()
			}
			if gotError != tt.wantErr {
				t.Errorf("requester.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			status := ""
			if got != nil {
				status = got.Status
			}

			if !reflect.DeepEqual(status, tt.wantStatus) {
				t.Errorf("requester.Do() = %v, wantStatus %v", status, tt.wantStatus)
			}
		})
	}
}

func TestRequestFromInternal(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
	}{
		{
			name: "should_copy_and_return_request_with_body",
			args: args{
				req: NewRequestMockWithBody(),
			},
			want: "{id:1}",
		},
		{
			name: "should_copy_and_return_request_with_body_nil",
			args: args{
				req: NewRequestMock(),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := requestFromInternal(tt.args.req)
			gotError := ""
			if err != nil {
				gotError = err.Error()
			}

			if gotError != tt.wantErr {
				t.Errorf("requester.requestFromInternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			body := ""
			if got.Body != nil {
				bytes, _ := io.ReadAll(got.Body)
				body = string(bytes)
			}

			if tt.want != body {
				t.Errorf("requester.requestFromInternal() = %v, want %v", body, tt.want)
			}
		})
	}
}

func TestCloseResponseBody(t *testing.T) {
	server, req := NewRequestWithHTTPServerOKMock()
	defer server.Close()

	s, reqWithResAndCancel, cancel := NewRequestWithHTTPServerUnavailableAndCanceledContext()
	defer s.Close()
	defer cancel()

	type args struct {
		req   *http.Request
		close func(*http.Response)
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			name: "should_close_body_after_read",
			args: args{
				req: req,
				close: func(r *http.Response) {
					r.Body.Close()
				},
			},
			wantErr: "http: read on closed response body",
		},
		{
			name: "should_not_close_body_after_read",
			args: args{
				req:   req,
				close: func(_ *http.Response) {},
			},
			wantErr: "",
		},
		{
			name: "should_close_body_when_response_has_error_and_context_is_canceled",
			args: args{
				req:   reqWithResAndCancel,
				close: func(_ *http.Response) {},
			},
			wantErr: "http: read on closed response body",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New()
			got, _ := d.Do(tt.args.req)

			tt.args.close(got)

			_, err := io.ReadAll(got.Body)
			gotError := ""
			if err != nil {
				gotError = err.Error()
			}

			if !reflect.DeepEqual(gotError, tt.wantErr) {
				t.Errorf("requester.Do() error = %v, wantError %v", err, tt.wantErr)
			}
		})
	}
}
