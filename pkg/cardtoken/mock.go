package cardtoken

import (
	"time"
)

func mockCardToken() *Response {
	return &Response{
		ID:              "3d40b34eb41a6d0923e5bc545927c2e9",
		FirstSixDigits:  "503143",
		ExpirationMonth: 11,
		ExpirationYear:  2025,
		LastFourDigits:  "6351",
		Cardholder: &Cardholder{
			Identification: &Identification{
				Number: "70383868084",
				Type:   "CPF",
			},
			Name: "MASTER TEST",
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

func MockCardTokenRequest() Request {
	return Request{
		SiteID:          "Teste",
		CardNumber:      "5031433215406351",
		ExpirationMonth: "11",
		ExpirationYear:  "2025",
		SecurityCode:    "123",
		Cardholder: &Cardholder{
			Identification: &Identification{
				Type:   "CPF",
				Number: "70383868084",
			},
			Name: "MASTER TEST",
		},
	}
}

func parseDate(s string) *time.Time {
	d, _ := time.Parse(time.RFC3339, s)
	return &d
}
