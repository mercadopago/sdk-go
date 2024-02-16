package oauth

type Response struct {
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	PublicKey    string `json:"public_key"`
	TokenType    string `json:"token_type"`
	LiveMode     bool   `json:"live_mode"`
	ExpiresIn    int64  `json:"expires_in"`
}
