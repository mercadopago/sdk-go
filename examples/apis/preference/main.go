package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	at := ("TEST-4679935697572392-071411-a9722b82869609999cd91f0db60598f0-1273205088")
	c, err := config.New(at)
	if err != nil {
		fmt.Println(err)
		return
	}
	get(c)
	create(c)
	update(c)
	search(c)
}

func get(c *config.Config) {
	pmc := preference.NewClient(c)
	res, err := pmc.Get(context.Background(), "1273205088-538fecdc-6a0b-4353-a514-0c6dda633b9f")
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

func create(c *config.Config) {
	pmc := preference.NewClient(c)

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

	resJSON, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(resJSON))
}

func update(c *config.Config) {
	pmc := preference.NewClient(c)

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

	res, err := pmc.Update(context.Background(), "1273205088-538fecdc-6a0b-4353-a514-0c6dda633b9f", dto)
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

func search(c *config.Config) {
	pmc := preference.NewClient(c)

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
