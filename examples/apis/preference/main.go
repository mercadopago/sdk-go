package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/option"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	cdt, err := credential.New("TEST-4679935697572392-071411-a9722b82869609999cd91f0db60598f0-1273205088")
	if err != nil {
		fmt.Println(err)
		return
	}
	getPreference(cdt)
	createPreference(cdt)
	
}

func getPreference(cdt credential.Credential) {
	pmc := preference.NewClient(
		option.WithRetryMax(1),
		option.WithTimeout(1*time.Second),
	)
	res, err := pmc.Get(context.Background(), cdt, "1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4")
	if err != nil {
		fmt.Println(err)
		return
	}

	resJSON, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(resJSON))
}

func createPreference(cdt credential.Credential) {
	pmc := preference.NewClient()

	dto := preference.Request{
		Items: []preference.PreferenceItemRequest{
			{
				ID:          "123",
				Title:       "Title",
				UnitPrice: 	100,
				Quantity:    1,
				Description: "Description",
			},
		},
	}

	res, err := pmc.Create(context.Background(), cdt, dto)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.ID)
}
