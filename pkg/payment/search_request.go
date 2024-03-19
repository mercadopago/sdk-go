package payment

import (
	"strconv"
	"strings"
)

// SearchRequest is the helper structure to build search request.
// Filters field can receive a lot of paramaters. For details, see:
// https://www.mercadopago.com/developers/en/reference/payments/_payments_search/get.
type SearchRequest struct {
	Limit   int
	Offset  int
	Filters map[string]string
}

// GetParams creates map to build query parameters. Keys will be changed to lower case.
func (sr *SearchRequest) GetParams() map[string]string {
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
