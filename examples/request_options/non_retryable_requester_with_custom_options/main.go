package main

import (
	"fmt"
	"time"

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

	nonRetryableRequester := httpclient.New(httpclient.WithTimeout(time.Second * 5))

	res, err := pc.Create(
		request,
		httpclient.WithRequestRequester(nonRetryableRequester), // sdk will use that requester
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.ID)
}
