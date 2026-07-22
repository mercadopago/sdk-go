package httpclient

import (
	"net/http"
	"testing"
)

func TestSetPathParamsEscapesPathTraversal(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "https://api.mercadopago.com/v1/customers/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := setPathParams(req, map[string]string{"id": "../../applications/123"}); err != nil {
		t.Fatal(err)
	}

	if req.URL.EscapedPath() != "/v1/customers/..%2F..%2Fapplications%2F123" {
		t.Fatalf("unexpected path: %s", req.URL.Path)
	}
}
