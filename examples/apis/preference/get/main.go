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
	pref, err := client.Get(context.Background(), "{{ID_PREFERENCE}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pref)
}
