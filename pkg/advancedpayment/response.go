package advancedpayment

import "time"

// Response represents an advanced (split) payment resource returned by the MercadoPago
// Advanced Payments API. It is returned by [Client.Create], [Client.Get], and related methods.
//
// An advanced payment allows a marketplace to collect a single payment and distribute the
// funds among multiple sellers (disbursements).
type Response struct {
	// Disbursements is the list of receiver/amount pairs associated with this payment.
	Disbursements []DisbursementResponse `json:"disbursements"`

	// Payer contains identifying information about the buyer.
	Payer PayerResponse `json:"payer"`

	// DateCreated is the date when the advanced payment was created.
	DateCreated time.Time `json:"date_created"`

	// DateLastUpdated is the date of the last update to this payment.
	DateLastUpdated time.Time `json:"date_last_updated"`

	// ApplicationID is the MercadoPago application identifier.
	ApplicationID string `json:"application_id"`

	// ExternalReference is the integrator-provided identifier for reconciliation.
	ExternalReference string `json:"external_reference"`

	// Description is a human-readable description of the payment.
	Description string `json:"description"`

	// Status is the overall status of the advanced payment.
	Status string `json:"status"`

	// ID is the unique advanced payment identifier assigned by MercadoPago.
	ID int `json:"id"`

	// BinaryMode when true means the payment is approved or rejected immediately.
	BinaryMode bool `json:"binary_mode"`

	// Capture indicates whether the payment has been captured.
	Capture bool `json:"capture"`
}

// DisbursementResponse represents a single receiver within an advanced payment.
type DisbursementResponse struct {
	// ExternalReference is the integrator-provided identifier for this disbursement.
	ExternalReference string `json:"external_reference"`

	// MoneyReleaseDate is the scheduled date when funds are released to the seller.
	MoneyReleaseDate string `json:"money_release_date"`

	// Status is the current status of this disbursement.
	Status string `json:"status"`

	// StatusDetail provides additional detail about the disbursement status.
	StatusDetail string `json:"status_detail"`

	// ApplicationFee is the fee retained by the marketplace.
	ApplicationFee float64 `json:"application_fee"`

	// Amount is the amount disbursed to this collector.
	Amount float64 `json:"amount"`

	// CollectorID is the MercadoPago user identifier of the seller.
	CollectorID int `json:"collector_id"`

	// ID is the unique disbursement identifier.
	ID int `json:"id"`
}

// PayerResponse contains information about the buyer of the advanced payment.
type PayerResponse struct {
	// Email is the payer's email address.
	Email string `json:"email"`

	// ID is the MercadoPago user identifier of the payer.
	ID string `json:"id"`

	// Type is the payer type.
	Type string `json:"type"`
}
