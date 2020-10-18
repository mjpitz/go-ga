package gatypes

// Transaction encapsulates fields on a transaction event
type Transaction struct { // t=transaction
	TransactionID          string  `url:"ti,omitempty"`
	TransactionAffiliation string  `url:"ta,omitempty"`
	TransactionRevenue     float64 `url:"tr,omitempty"`
	TransactionShipping    float64 `url:"ts,omitempty"`
	TransactionTax         float64 `url:"tt,omitempty"`
}
