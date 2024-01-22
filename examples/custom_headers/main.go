package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/header"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	cdt, err := credential.New("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")
	if err != nil {
		fmt.Println(err)
		return
	}

	pmc := paymentmethod.NewClient()

	ch := http.Header{}
	ch.Add("X-Idempotency-Key", "123999")
	ch.Add("Some-Key", "some_value")

	ctx := context.Background()
	// this will return a child context decorated with your custom headers.
	// They will be forwarded wherever this context is used in any of the APIs.
	ctx = header.Context(ctx, ch)
	res, err := pmc.List(
		ctx,
		cdt,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
