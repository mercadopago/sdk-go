package preapprovalplan

import "strings"

// SearchRequest contains filters accepted in search.
// Filters field can receive a lot of parameters. For details, see:
// https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_plan_id/get
type SearchRequest struct {
	Filters map[string]string

	Limit  string
	Offset string
}

// SetDefaults sets values for limit and offset when not sent.
func (s *SearchRequest) SetDefaults() {
	if len(s.Filters) == 0 {
		s.Filters = make(map[string]string, 2)
	} else {
		for k, v := range s.Filters {
			delete(s.Filters, k)
			s.Filters[strings.ToLower(k)] = v
		}
	}

	if _, ok := s.Filters["limit"]; !ok {
		limit := "30"
		if s.Limit != "" {
			limit = s.Limit
		}
		s.Filters["limit"] = limit
	}
	if _, ok := s.Filters["offset"]; !ok {
		offset := "0"
		if s.Offset != "" {
			offset = s.Offset
		}
		s.Filters["offset"] = offset
	}
}
