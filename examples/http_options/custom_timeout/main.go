package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/option"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	cdt, err := credential.New("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")
	if err != nil {
		fmt.Println(err)
		return
	}

	pmc := paymentmethod.NewClient(option.WithTimeout(time.Second * 3))
	res, err := pmc.List(context.Background(), cdt)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}