package main

import (
	"fmt"
	"time"

	"github.com/mercadopago/sdk-go/pkg/mp"
	"github.com/mercadopago/sdk-go/pkg/mp/rest"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	rc := mp.NewRestClient("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

	pc := payment.NewClient(rc)

	request := payment.Request{
		TransactionAmount: 1.5,
		PaymentMethodID:   "pix",
		Description:       "meu pagamento",
		Payer: &payment.PayerRequest{
			Email: "fhashfadsuhfdafasdfasfashfda@testuser.com",
		},
	}

	opts := []rest.Option{
		rest.WithTimeout(time.Nanosecond * 1), // request timeout will be 10 seconds
	}

	res, err := pc.Create(request, opts...)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.ID)
}
