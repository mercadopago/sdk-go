package user

type Response struct {
	ID        int64  `json:"id"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CountryID string `json:"country_id"`
	Email     string `json:"email"`
	SiteID    string `json:"site_id"`
}
