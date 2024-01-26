package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/option"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	at := "TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800"

	c, err := config.New(
		at,
		option.WithCorporationID("abc"),
		option.WithIntegratorID("def"),
		option.WithPlatformID("ghi"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	pmc := paymentmethod.NewClient(c)
	res, err := pmc.List(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
