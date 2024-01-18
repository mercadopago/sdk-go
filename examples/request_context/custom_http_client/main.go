package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
	"github.com/mercadopago/sdk-go/pkg/request"
)

func main() {
	cdt, err := credential.New("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")
	if err != nil {
		fmt.Println(err)
		return
	}

	pmc := paymentmethod.NewClient()

	customClient := &http.Client{}

	ctx := context.Background()
	ctx = request.WithCustomClient(ctx, customClient) // sdk will use that http client
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
