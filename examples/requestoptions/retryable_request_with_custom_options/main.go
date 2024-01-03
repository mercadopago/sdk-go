package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
	"github.com/mercadopago/sdk-go/pkg/mp"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	mp.SetAccessToken("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

	pc := payment.NewClient()

	request := payment.Request{
		TransactionAmount: 1.5,
		PaymentMethodID:   "pix",
		Description:       "my payment",
		Payer: &payment.PayerRequest{
			Email: "gabs@testuser.com",
		},
	}

	retryMax := 5

	backoffStrategy := func(attempt int) time.Duration {
		// your retry implementation, for more information see: httpclient.BackoffFunc
		return time.Duration(1)
	}
	checkRetry := func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		// your retry stop condition, for more information see: httpclient.CheckRetryFunc
		return false, nil
	}

	options := []httpclient.OptionRetryable{
		httpclient.WithTimeout(time.Nanosecond * 1),
		httpclient.WithBackoffStrategy(backoffStrategy),
		httpclient.WithRetryPolicy(checkRetry),
	}

	customRetryableRequester := httpclient.NewRetryable(retryMax, options...)

	opts := []httpclient.RequestOption{
		httpclient.WithRequestRequester(customRetryableRequester), // http client will use these custom headers
	}

	res, err := pc.Create(request, opts...)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.ID)
}
