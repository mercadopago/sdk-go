package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := preference.NewClient(cfg)

	preferenceID := "123"

	pref, err := client.Get(context.Background(), preferenceID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pref)
}
