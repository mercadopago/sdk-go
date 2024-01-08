package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mercadopago/sdk-go/pkg/api"
	"github.com/mercadopago/sdk-go/pkg/api/paymentmethod"
	"github.com/mercadopago/sdk-go/pkg/httpclient"
	"github.com/mercadopago/sdk-go/pkg/mp"
)

func main() {
	mp.SetAccessToken("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

	pmc := paymentmethod.NewClient()

	var backoffStrategy httpclient.BackoffFunc = func(attempt int) time.Duration {
		// your retry implementation, for more information see: httpclient.BackoffFunc
		return time.Duration(1)
	}

	var checkRetry httpclient.CheckRetryFunc = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		// your retry stop condition, for more information see: httpclient.CheckRetryFunc
		return false, nil
	}

	retryableRequester := httpclient.NewRetryable(
		httpclient.WithRetryMax(5),
		httpclient.WithTimeout(time.Second*1),
		httpclient.WithBackoffStrategy(backoffStrategy),
		httpclient.WithRetryPolicy(checkRetry),
	)

	res, err := pmc.List(
		api.WithRequester(retryableRequester), // sdk will use that requester
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
