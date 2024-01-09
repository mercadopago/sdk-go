package main

import (
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/api"
	"github.com/mercadopago/sdk-go/pkg/api/paymentmethod"
	"github.com/mercadopago/sdk-go/pkg/mp"
)

func main() {
	mp.SetAccessToken("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

	pmc := paymentmethod.NewClient()

	ch := http.Header{}
	ch.Add("X-Idempotency-Key", "123999")
	ch.Add("Some-Key", "some_value")
	res, err := pmc.List(
		api.WithCustomHeaders(ch), // http client will use these custom headers
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
