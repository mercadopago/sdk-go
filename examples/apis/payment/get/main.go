package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	at := "TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800"
	c, err := config.New(at)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := payment.NewClient(c)
	result, err := client.Get(context.Background(), 123)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
