package oauth

// Request represents the body payload sent to the MercadoPago OAuth API when creating
// new credentials via [Client.Create]. It exchanges an authorization code for an access
// token and refresh token using the "authorization_code" grant type.
type Request struct {
	// GrantType is the OAuth grant type. Set to "authorization_code" for credential creation.
	GrantType string `json:"grant_type,omitempty"`

	// ClientSecret is the application's access token used to authenticate the OAuth request.
	ClientSecret string `json:"client_secret,omitempty"`

	// Code is the authorization code received from the OAuth authorization flow redirect.
	Code string `json:"code,omitempty"`

	// RedirectURI is the callback URL that was used when generating the authorization URL.
	// It must match exactly for the token exchange to succeed.
	RedirectURI string `json:"redirect_uri,omitempty"`
}

// RefreshTokenRequest represents the body payload sent to the MercadoPago OAuth API when
// refreshing expired credentials via [Client.Refresh]. It uses the "refresh_token" grant
// type to obtain a new access token without requiring the seller to re-authorize.
type RefreshTokenRequest struct {
	// GrantType is the OAuth grant type. Set to "refresh_token" for credential refresh.
	GrantType string `json:"grant_type,omitempty"`

	// ClientSecret is the application's access token used to authenticate the OAuth request.
	ClientSecret string `json:"client_secret,omitempty"`

	// RefreshToken is the refresh token obtained from a previous [Client.Create] or [Client.Refresh] call.
	RefreshToken string `json:"refresh_token,omitempty"`
}
