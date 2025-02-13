package main

import (
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"
	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}
}
