package create

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/cardtoken"
	"github.com/mercadopago/sdk-go/pkg/config"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := cardtoken.NewClient(cfg)

	result, err := client.Create(context.Background(), cardtoken.Request{})
	if err != nil {
		return
	}

	fmt.Println(result)
}
