package gatypes

// Item encapsulates information about an item event.
type Item struct { // t=item
	ItemName     string  `url:"in,omitempty" json:"item_name,omitempty"`
	ItemPrice    float64 `url:"ip,omitempty" json:"item_price,omitempty"`
	ItemQuantity int64   `url:"iq,omitempty" json:"item_quantity,omitempty"`
	ItemCode     string  `url:"ic,omitempty" json:"item_code,omitempty"`
	ItemCategory string  `url:"iv,omitempty" json:"item_category,omitempty"`
}
