// Package main shows how to create an Order using PSE (Pagos Seguros en Línea — Colombia).
//
// PSE is Colombia's standard online bank-transfer payment method. The integrator
// initiates the transaction with PaymentMethod.ID = "pse" and Type = "bank_transfer",
// and must specify the buyer's bank via FinancialInstitution.
//
// Required PSE-specific fields:
//   - PaymentMethod.ID = "pse" (fixed)
//   - PaymentMethod.Type = "bank_transfer" (fixed)
//   - PaymentMethod.FinancialInstitution = PSE bank code (see table below)
//   - Currency = "COP" (Colombia only)
//   - Payer.EntityType = "individual" or "association"
//   - Payer.Identification.Type = "CC", "NIT", etc. (Colombian doc type)
//   - AdditionalInfo.PayerIPAddress (required by the risk engine)
//   - Config.Online.CallbackURL (URL the bank redirects to after authorization)
//
// Most-used PSE bank codes (full catalog via MP API):
//
//	Bancolombia ........... 1007
//	Davivienda ............ 1051
//	Banco de Bogotá ....... 1001
//	BBVA Colombia ......... 1013
//	Banco Popular ......... 1002
//	Scotiabank Colpatria .. 1019
//
// Reference: https://www.mercadopago.com.co/developers/es/docs
package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"
	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)
	request := order.Request{
		Type:              "online",
		ProcessingMode:    "automatic",
		TotalAmount:       "50000.00",
		Currency:          "COP",
		ExternalReference: "ref_pse_12345",
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "50000.00",
					PaymentMethod: &order.PaymentMethodRequest{
						ID:                   "pse",
						Type:                 "bank_transfer",
						FinancialInstitution: "1007", // Bancolombia
					},
				},
			},
		},
		// Payer: entity_type + Colombian identification (CC/NIT) required for PSE.
		Payer: &order.PayerRequest{
			Email:      "{{PAYER_EMAIL}}",
			FirstName:  "{{FIRST_NAME}}",
			LastName:   "{{LAST_NAME}}",
			EntityType: "individual",
			Identification: &order.IdentificationRequest{
				Type:   "CC",
				Number: "{{PAYER_DOC_NUMBER}}",
			},
		},
		// additional_info.payer.ip_address — required by MP's risk engine for PSE.
		AdditionalInfo: &order.AdditionalInfoRequest{
			PayerIPAddress: "{{CLIENT_IP}}",
		},
		// callback_url — where the bank redirects the buyer after authorization.
		Config: &order.ConfigRequest{
			Online: &order.OnlineConfigRequest{
				CallbackURL: "{{CALLBACK_URL}}",
			},
		},
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
