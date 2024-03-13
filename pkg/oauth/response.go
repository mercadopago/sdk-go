package oauth

// Response represents credential information for an Oauth authorization
type Response struct {
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	PublicKey    string `json:"public_key"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	LiveMode     bool   `json:"live_mode"`
}
