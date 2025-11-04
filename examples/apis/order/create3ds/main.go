package main

import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/mercadopago/sdk-go/pkg/config"
    "github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
    accessToken := "{{ACCESS_TOKEN}}"

    cfg, err := config.New(accessToken)
    if err != nil {
        fmt.Println(err)
        return
    }

    client := order.NewClient(cfg)

    request := order.Request{
        Type:              "online",
        TotalAmount:       "{{TOTAL_AMOUNT}}",       // ex.: "150.00"
        ExternalReference: "{{EXTERNAL_REFERENCE}}", // ex.: "ext_ref_3ds_123"
        Payer: &order.PayerRequest{
            Email: "{{EMAIL}}",
        },
        Transactions: &order.TransactionRequest{
            Payments: []order.PaymentRequest{
                {
                    Amount: "{{AMOUNT}}", // ex.: "150.00"
                    PaymentMethod: &order.PaymentMethodRequest{
                        ID:           "{{PAYMENT_METHOD_ID}}", // ex.: "master"
                        Token:        "{{CARD_TOKEN}}",
                        Type:         "{{TYPE}}", // ex.: "credit_card"
                        Installments: 1,
                    },
                },
            },
        },
        Config: &order.ConfigRequest{
            Online: &order.OnlineConfigRequest{
                TransactionSecurity: &order.TransactionSecurityRequest{
                    Validation:     "on_fraud_risk",
                    LiabilityShift: "required",
                },
                // Opcional: configure URLs de retorno/callback conforme sua aplicação
                // SuccessURL:  "https://example.com/success",
                // FailureURL:  "https://example.com/failure",
                // PendingURL:  "https://example.com/pending",
                // CallbackURL: "https://example.com/webhook",
            },
        },
    }

    resource, err := client.Create(context.Background(), request)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("== CRIAR ORDER (3DS) ==")
    fmt.Println("✓ ORDER CRIADA COM SUCESSO!")
    fmt.Printf("ID: %s\n", resource.ID)
    fmt.Printf("Status: %s (%s)\n", resource.Status, resource.StatusDetail)

    if len(resource.Transactions.Payments) > 0 {
        p := resource.Transactions.Payments[0]
        fmt.Println("\n-- Pagamento --")
        fmt.Printf("Payment ID: %s\n", p.ID)
        fmt.Printf("Status: %s (%s)\n", p.Status, p.StatusDetail)
        if p.PaymentMethod.TransactionSecurity != nil {
            ts := p.PaymentMethod.TransactionSecurity
            fmt.Println("\n-- 3DS / Transaction Security --")
            fmt.Printf("Validation: %s\n", ts.Validation)
            fmt.Printf("Liability Shift: %s\n", ts.LiabilityShift)
            fmt.Printf("Type: %s | Status: %s\n", ts.Type, ts.Status)
            if ts.URL != "" {
                fmt.Printf("Challenge URL: %s\n", ts.URL)
                fmt.Println("Abra essa URL no navegador/iframe para conduzir o challenge 3DS.")
            }
        }
    }

    fmt.Println("\n-- Resposta JSON Completa --")
    pretty, _ := json.MarshalIndent(resource, "", "  ")
    fmt.Println(string(pretty))
}


