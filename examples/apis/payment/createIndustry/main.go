package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	accessToken := "APP_USR-874202490252970-100714-e890db6519b0dceb4ef24ef41ed816e4-2021490138"
	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println("Config error:", err)
		return
	}

	client := payment.NewClient(cfg)

	request := payment.Request{
		TransactionAmount:   150.00,
		Description:         "Teste dados da industria",
		PaymentMethodID:     "master",
		Token:               "1b7821136a3a321512fec7defea47f89",
		Installments:        1,
		BinaryMode:          false,
		Capture:             true,
		ExternalReference:   "Pedido01",
		NotificationURL:     "http://requestbin.fullcontact.com/1ogudgk1",
		StatementDescriptor: "LOJA 123",
		Payer: &payment.PayerRequest{
			FirstName: "Guilherme",
			LastName:  "Santos",
			Email:     "teste_dados_industria@testuser.com",
			Identification: &payment.IdentificationRequest{
				Type:   "CPF",
				Number: "12345678909",
			},
		},
		Metadata: map[string]interface{}{
			"order_number": "order_01",
		},
		/*ForwardData: &payment.ForwardDataRequest{
			SubMerchant: &payment.SubMerchantRequest{
				SubMerchantID:     "1234183712",
				MCC:               "5462",
				Country:           "BRA",
				AddressDoorNumber: 123,
				Zip:               "222222",
				DocumentNumber:    "222222222222222",
				City:              "SÃO PAULO",
				LegalName:         "LOJINHA DO ZÉ",
				AddressStreet:     "RUA A",
				RegionCodeISO:     "BR-MG",
				RegionCode:        "BR",
				DocumentType:      "CNPJ",
				Phone:             "123123123",
				URL:               "www.rappi.com.br",
			},
		},*/
		/*AdditionalInfo: &payment.AdditionalInfoRequest{
			Items: []payment.ItemRequest{
				{
					ID:          "1941",
					Title:       "25/08/2022 | Pista Inteira5 lote - GREEN VALLEY GRAMADO 2022",
					Description: "25/08/2022 | Pista Inteira5 lote - GREEN VALLEY GRAMADO 2022",
					CategoryID:  "Tickets",
					Quantity:    1,
					UnitPrice:   1000.00,
					EventDate:   "2019-10-25T19:30:00.000-03:00",
					PictureURL:  "{{url_image}}",
					Type:        "my_items_type",
					Warranty:    true,
					Category_Descriptor: &payment.CategoryDescriptorRequest{
						Passenger: &payment.PassengerRequest{
							FirstName: "Guest Nome",
							LastName:  "Guest Sobrenome",
							Identification: &payment.IdentificationRequest{
								Type:   "DNI",
								Number: "012345678",
							},
						},
						Route: &payment.RouteRequest{
							ArrivalDateTime:   "2019-12-25T19:30:00.000-03:00",
							Company:           "Companhia",
							Departure:         "Osasco",
							DepartureDateTime: "2022-03-12T12:58:41.425-04:00",
							Destination:       "Sao Paulo",
						},
					},
				},
			},
			Payer: &payment.AdditionalInfoPayerRequest{
				FirstName:             "Nome",
				LastName:              "Sobrenome",
				IsPrimeUser:           true,
				IsFirstPurchaseOnline: true,
				LastPurchase:          "2019-10-25T19:30:00.000-03:00",
				Phone: &payment.AdditionalInfoPayerPhoneRequest{
					AreaCode: "11",
					Number:   "987654321",
				},
				Address: &payment.AdditionalInfoPayerAddressRequest{
					StreetName:   "Av. das Nações Unidas",
					StreetNumber: "3003",
					ZipCode:      "206233-2002",
				},
				RegistrationDate:   "2020-08-06T09:25:04.000-03:00",
				AuthenticationType: "Gmail",
			},
			Shipments: &payment.ShipmentsRequest{
				LocalPickup:     true,
				ExpressShipment: true,
				ReceiverAddress: &payment.ReceiverAddressRequest{
					ZipCode:      "306233-2003",
					StreetName:   "Av. das Nações Unidas",
					StreetNumber: "3003",
					Floor:        "5",
					Apartment:    "502",
					StateName:    "DF",
					CityName:     "Bogota",
				},
			},
		},*/
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println("Payment error:", err)
		return
	}

	fmt.Println("Payment ID:", resource.ID)
}
