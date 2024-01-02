package payment

// Filters is the filters to search for payments.
type Filters struct {
	// Sort is a field used to sort a list of payments.
	// The sort can be done by the following attributes:
	//	- "date_approved"
	//	- "date_created"
	//	- "date_last_updated"
	//	- "id"
	//	- "money_release_date"
	Sort string

	// Criteria is a field used to define the order of the result list.
	// Can be "asc" or "desc".
	Criteria string

	// ExternalReference is an external reference of the payment.
	// It can be, for example, a hashcode from the Central Bank, working as an identifier of the transaction origin.
	ExternalReference string

	// Range is a field used to define the range of the search.
	// The Range can be related to the following attributes:
	//	- "date_created"
	//	- "date_last_updated"
	//	- "date_approved"
	//	- "money_release_date"
	// If not informed, it uses "date_created" by default.
	Range string

	// BeginDate is a field used to define the start of the search interval.
	// Its format can be a relative date - "NOW-XDAYS", "NOW-XMONTHS" - or an absolute date - ISO8601.
	// If not informed, it uses "NOW-3MONTHS" by default.
	BeginDate string

	// EndDate is a field used to define the end of the search interval.
	// Its format can be a relative date - "NOW-XDAYS", "NOW-XMONTHS" - or an absolute date - ISO8601.
	// If not informed, it uses "NOW" by default.
	EndDate string
}
