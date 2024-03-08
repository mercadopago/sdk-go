package preapproval

import (
	"strconv"
	"strings"
)

// SearchRequest contains filters accepted in search.
// Filters field can receive a lot of parameters. For details, see:
// https://www.mercadopago.com.br/developers/pt/reference/subscriptions/_preapproval_search/get
type SearchRequest struct {
	Filters map[string]string

	Limit  int
	Offset int
}

// GetParams creates map to build query parameters. Keys will be changed to lower case.
func (sr SearchRequest) GetParams() map[string]string {
	params := map[string]string{}
	for k, v := range sr.Filters {
		key := strings.ToLower(k)
		params[key] = v
	}

	if sr.Limit == 0 {
		sr.Limit = 30
	}
	params["limit"] = strconv.Itoa(sr.Limit)
	params["offset"] = strconv.Itoa(sr.Offset)

	return params
}
