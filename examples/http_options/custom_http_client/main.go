package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

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

	proxyURL, _ := url.Parse("http://someurl")
	customClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	pmc := paymentmethod.NewClient(option.WithCustomClient(customClient))
	res, err := pmc.List(context.Background(), cdt)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
