package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mercadopago/sdk-go/pkg/api/paymentmethod"
	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/httpclient"
)

var (
	customBackoffStrategy httpclient.BackoffFunc = func(attempt int) time.Duration {
		// your backoff strategy
		return time.Duration(1)
	}

	customCheckRetry httpclient.CheckRetryFunc = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		// your logic to define continue or stop retry condition
		return false, nil
	}
)

func main() {
	cdt, err := credential.New("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")
	if err != nil {
		fmt.Println(err)
		return
	}

	pmc := paymentmethod.NewClient(
		httpclient.WithRetryMax(5),
		httpclient.WithBackoffStrategy(customBackoffStrategy),
		httpclient.WithRetryPolicy(customCheckRetry),
		httpclient.WithTimeout(time.Millisecond*1000),
	)
	res, err := pmc.List(cdt)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
