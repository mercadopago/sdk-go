package cardtoken

import (
	"time"
)

func mockCardToken() *Response {
	return &Response{
		Id:              "3d40b34eb41a6d0923e5bc545927c2e9",
		FirstSixDigits:  "123456",
		ExpirationMonth: 11,
		ExpirationYear:  2030,
		LastFourDigits:  "1234",
		Cardholder: Cardholder{
			Identification: Identification{
				Number: "12345678901",
				Type:   "CPF",
			},
			Name: "APRO",
		},
		Status:             "active",
		DateCreated:        parseDate("2024-02-08T09:05:42.725-04:00"),
		DateLastUpdated:    parseDate("2024-02-08T09:05:42.725-04:00"),
		DateDue:            parseDate("2024-02-16T09:05:42.725-04:00"),
		LuhnValidation:     true,
		LiveMode:           false,
		CardNumberLength:   16,
		SecurityCodeLength: 3,
	}
}

func parseDate(s string) time.Time {
	d, _ := time.Parse(time.RFC3339, s)
	return d
}
