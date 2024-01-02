package main

import (
	"fmt"
	"net/http"

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

	ch := http.Header{}
	ch.Add("X-Idempotency-Key", "123")
	ch.Add("Some-Key", "some_value")
	opts := []rest.Option{
		rest.WithCustomHeaders(ch), // rest client will use these custom headers
	}

	res, err := pc.Create(request, opts...)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.ID)
}
