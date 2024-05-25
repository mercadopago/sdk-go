package oauth

// Request represents credential information to perform a create credential request.
type Request struct {
	GrantType    string `json:"grant_type,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Code         string `json:"code,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

// RefreshTokenRequest represents credential information to perform a refresh credential request.
type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
