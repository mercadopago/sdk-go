package oauth

// Response represents the OAuth credentials returned by the MercadoPago OAuth API
// after a successful [Client.Create] or [Client.Refresh] call. It contains the access
// token needed to make API calls on behalf of a seller, along with its expiration and scope.
type Response struct {
	// AccessToken is the Bearer token used to authenticate API requests on behalf of the seller.
	AccessToken string `json:"access_token"`

	// Scope defines the permissions granted by the seller (e.g., "read", "write", "offline_access").
	Scope string `json:"scope"`

	// RefreshToken is used to obtain a new access token when the current one expires,
	// without requiring the seller to re-authorize. Pass it to [Client.Refresh].
	RefreshToken string `json:"refresh_token"`

	// PublicKey is the public key associated with the seller's credentials, used for
	// client-side tokenization in frontend integrations.
	PublicKey string `json:"public_key"`

	// TokenType is the type of the access token, typically "Bearer".
	TokenType string `json:"token_type"`

	// ExpiresIn is the number of seconds until the access token expires.
	ExpiresIn int `json:"expires_in"`

	// LiveMode indicates whether the credentials operate in production (true) or sandbox (false).
	LiveMode bool `json:"live_mode"`
}
