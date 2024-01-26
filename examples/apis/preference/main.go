package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	cdt, err := credential.New("TEST-4679935697572392-071411-a9722b82869609999cd91f0db60598f0-1273205088")
	if err != nil {
		fmt.Println(err)
		return
	}
	get(*cdt)
	create(*cdt)
	update(*cdt)
	search(*cdt)
	
}

func get(cdt credential.Credential) {
	pmc := preference.NewClient(&cdt)
	res, err := pmc.Get(context.Background(), "1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4")
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

func create(cdt credential.Credential) {
	pmc := preference.NewClient(&cdt)

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

	res, err := pmc.Create(context.Background(), dto)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.ID)
}

func update(cdt credential.Credential) {
	pmc := preference.NewClient(&cdt)

	dto := preference.Request{
		Items: []preference.PreferenceItemRequest{
			{
				ID:          "123",
				Title:       "updating Title",
				UnitPrice: 	10,
				Quantity:    3,
				Description: "updating Description",
			},
		},
	}

	res, err := pmc.Update(context.Background(), "1273205088-13736a46-a3e0-45bb-b610-2cef417f8da4", dto)
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

func search(cdt credential.Credential) {
	pmc := preference.NewClient(&cdt)

	filters := preference.SearchRequest{
		Limit: 10,
		Offset: 10,
	}

	res, err := pmc.Search(context.Background(), filters)
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
