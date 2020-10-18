package gatypes

// Item encapsulates information about an item event.
type Item struct { // t=item
	TransactionID string  `url:"ti,omitempty"`
	ItemName      string  `url:"in,omitempty"`
	ItemPrice     float64 `url:"ip,omitempty"`
	ItemQuantity  int64   `url:"iq,omitempty"`
	ItemCode      string  `url:"ic,omitempty"`
	ItemCategory  string  `url:"iv,omitempty"`
}
