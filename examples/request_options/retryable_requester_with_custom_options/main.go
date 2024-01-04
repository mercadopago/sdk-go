package main

// func main() {
// 	mp.SetAccessToken("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

// 	pc := payment.NewClient()

// 	request := payment.Request{
// 		TransactionAmount: 1.5,
// 		PaymentMethodID:   "pix",
// 		Description:       "my payment",
// 		Payer: &payment.PayerRequest{
// 			Email: "gabs@testuser.com",
// 		},
// 	}

// 	retryMax := 5

// 	backoffStrategy := func(attempt int) time.Duration {
// 		// your retry implementation, for more information see: httpclient.BackoffFunc
// 		return time.Duration(1)
// 	}
// 	checkRetry := func(ctx context.Context, resp *http.Response, err error) (bool, error) {
// 		// your retry stop condition, for more information see: httpclient.CheckRetryFunc
// 		return false, nil
// 	}

// 	retryableRequester := httpclient.NewRetryable(
// 		retryMax,
// 		httpclient.WithTimeout(time.Nanosecond*1),
// 		httpclient.WithBackoffStrategy(backoffStrategy),
// 		httpclient.WithRetryPolicy(checkRetry),
// 	)

// 	opts := []httpclient.RequestOption{
// 		httpclient.WithRequestRequester(retryableRequester), // sdk will use that requester
// 	}

// 	res, err := pc.Create(request, opts...)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(res.ID)
// }
