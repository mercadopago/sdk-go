package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
	"github.com/mercadopago/sdk-go/pkg/mp"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	mp.SetAccessToken("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

	proxyURL, _ := url.Parse("http://someurl")
	customClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	pmc := paymentmethod.NewClient(httpclient.WithAPIRequester(customClient))
	res, err := pmc.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
