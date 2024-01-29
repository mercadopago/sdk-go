package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/option"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"
	httpClient := &http.Client{Timeout: time.Second * 5}

	cfg, err := config.New(accessToken, option.WithHTTTPClient(httpClient))
	if err != nil {
		fmt.Println(err)
		return
	}

	client := paymentmethod.NewClient(cfg)
	paymentMethods, err := client.List(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range paymentMethods {
		fmt.Println(v)
	}
}
