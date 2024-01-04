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

// 	requester := httpclient.New()

// 	res, err := pc.Create(
// 		request,
// 		httpclient.WithRequestRequester(requester), // sdk will use that requester
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(res.ID)
// }
