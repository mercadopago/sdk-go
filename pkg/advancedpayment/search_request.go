package advancedpayment

import (
	"strconv"
	"strings"
)

// SearchRequest holds the parameters for searching advanced payments via [Client.Search].
type SearchRequest struct {
	// Limit is the maximum number of results per page. Defaults to 30 if zero.
	Limit int

	// Offset is the number of results to skip for pagination.
	Offset int

	// Filters contains additional query parameters (e.g., "status", "external_reference").
	// Keys are lowercased before being sent to the API.
	Filters map[string]string
}

// GetParams converts the SearchRequest into query parameters for the MercadoPago API.
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
