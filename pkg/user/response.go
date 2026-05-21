package user

// Response represents the authenticated MercadoPago user's account information
// as returned by [Client.Get].
type Response struct {
	// Nickname is the user's MercadoPago display name or username.
	Nickname string `json:"nickname"`

	// FirstName is the user's first name.
	FirstName string `json:"first_name"`

	// LastName is the user's last name.
	LastName string `json:"last_name"`

	// CountryID is the ISO 3166-1 alpha-2 country code of the user's account (e.g., "BR", "AR", "MX").
	CountryID string `json:"country_id"`

	// Email is the user's email address.
	Email string `json:"email"`

	// SiteID is the MercadoPago site identifier for the user's account (e.g., "MLB", "MLA", "MLM").
	SiteID string `json:"site_id"`

	// ID is the user's unique MercadoPago numeric identifier.
	ID int `json:"id"`
}
