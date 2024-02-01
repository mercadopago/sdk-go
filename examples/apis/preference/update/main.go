package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	at := "{{ACCESS_TOKEN}}"
	c, err := config.New(at)
	if err != nil {
		fmt.Println(err)
		return
	}

	request := preference.Request{
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

	client := preference.NewClient(c)
	res, err := client.Update(context.Background(), "id", request)
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
