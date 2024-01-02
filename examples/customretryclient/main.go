package main

import (
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/mp"
	"github.com/mercadopago/sdk-go/pkg/mp/rest"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

type customRetryClient struct{}

func (*customRetryClient) Retry(req *http.Request, httpClient *http.Client, opts ...rest.Option) (*http.Response, error) {
	// some retry implementation
	return nil, nil
}

func main() {
	rc := mp.NewRestClient("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

	customRetryClient := &customRetryClient{}
	mp.SetCustomRetryClient(customRetryClient)

	pmc := paymentmethod.NewClient(rc)
	res, err := pmc.List()
	if err != nil {
		panic(err)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
