package oauth

type Request struct {
	GrantType    string `json:"grant_type,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	Code         string `json:"code,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
