package gatypes

// Transaction encapsulates fields on a transaction event
type Transaction struct { // t=transaction
	TransactionAffiliation string  `url:"ta,omitempty" json:"transaction_affiliation,omitempty"`
	TransactionRevenue     float64 `url:"tr,omitempty" json:"transaction_revenue,omitempty"`
	TransactionShipping    float64 `url:"ts,omitempty" json:"transaction_shipping,omitempty"`
	TransactionTax         float64 `url:"tt,omitempty" json:"transaction_tax,omitempty"`
}
