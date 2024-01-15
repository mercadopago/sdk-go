package main

import (
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/api"
	"github.com/mercadopago/sdk-go/pkg/api/paymentmethod"
	"github.com/mercadopago/sdk-go/pkg/credential"
)

type myRequester struct{}

func (*myRequester) Do(req *http.Request) (*http.Response, error) {
	// my own Do logic
	return nil, nil
}

func main() {
	cdt, err := credential.New("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")
	if err != nil {
		fmt.Println(err)
		return
	}

	pmc := paymentmethod.NewClient()

	// can be a http.Client from standard library:
	// standardLibClient := &http.Client{}
	// or can be a custom requester
	myOwnRequester := &myRequester{}

	res, err := pmc.List(
		cdt,
		api.WithRequester(myOwnRequester), // sdk will use that requester
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
