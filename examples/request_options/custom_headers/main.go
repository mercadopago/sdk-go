package main

import (
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
	"github.com/mercadopago/sdk-go/pkg/mp"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

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

	ch := http.Header{}
	ch.Add("X-Idempotency-Key", "123999")
	ch.Add("Some-Key", "some_value")
	opts := []httpclient.RequestOption{
		httpclient.WithCustomHeaders(ch), // http client will use these custom headers
	}

	res, err := pc.Create(request, opts...)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.ID)
	}
}
