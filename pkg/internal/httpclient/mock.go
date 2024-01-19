package httpclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func Mock(response string) (*httptest.Server, error) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/fixedvalue" {
			panic(fmt.Sprintf("Expected to request '/fixedvalue', got: %s", r.URL.Path))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	return server, nil
}
