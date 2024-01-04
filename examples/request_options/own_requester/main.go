package main

import (
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
	"github.com/mercadopago/sdk-go/pkg/mp"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

type myRequester struct{}

func (*myRequester) Do(req *http.Request) (*http.Response, error) {
	// my own Do logic
	return nil, nil
}

func main() {
	mp.SetAccessToken("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

	pc := payment.NewClient()

	request := payment.Request{
		TransactionAmount: 1.5,
		PaymentMethodID:   "pix",
		Description:       "my payment",
		Payer: &payment.PayerRequest{
			Email: "gabs@testuser.com",
		},
	}

	// can be a http.Client from standard library:
	// standardLibClient := &http.Client{}
	// or can be a custom requester
	myOwnRequester := &myRequester{}

	opts := []httpclient.RequestOption{
		httpclient.WithRequestRequester(myOwnRequester), // sdk will use that requester
	}

	res, err := pc.Create(request, opts...)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.ID)
	}
}
