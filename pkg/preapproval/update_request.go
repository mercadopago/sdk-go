package preapproval

import "time"

// UpdateRequest represents the body sent to modify an existing pre-approval (subscription)
// in the MercadoPago Subscriptions API. Only non-zero fields are included in the JSON
// payload, so callers only need to set the fields they want to change.
//
// Use this struct with [Client.Update].
type UpdateRequest struct {
	// AutoRecurring contains the recurring billing fields to update, such as amount or dates.
	AutoRecurring *AutoRecurringUpdateRequest `json:"auto_recurring,omitempty"`

	// CardTokenID is a new card token to replace the current payment card for the subscription.
	CardTokenID string `json:"card_token_id,omitempty"`
	// PayerEmail is the updated email address of the subscriber.
	PayerEmail string `json:"payer_email,omitempty"`
	// BackURL is the updated redirect URL after the subscription flow completes.
	BackURL string `json:"back_url,omitempty"`
	// Reason is the updated description of the subscription shown to the payer.
	Reason string `json:"reason,omitempty"`
	// ExternalReference is the updated external reference value for correlating with
	// an entity in the integrator's system.
	ExternalReference string `json:"external_reference,omitempty"`
	// Status is the updated subscription status (e.g., "paused", "cancelled").
	Status string `json:"status,omitempty"`
}

// AutoRecurringUpdateRequest represents the subset of recurring billing fields that can be
// modified on an existing pre-approval (subscription). Unlike [AutoRecurringRequest], it
// does not allow changing the currency or frequency because those are immutable after creation.
type AutoRecurringUpdateRequest struct {
	// StartDate is the updated date when recurring charges begin.
	StartDate *time.Time `json:"start_date,omitempty"`
	// EndDate is the updated date when recurring charges end.
	EndDate *time.Time `json:"end_date,omitempty"`

	// TransactionAmount is the updated amount charged per billing cycle.
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
}
