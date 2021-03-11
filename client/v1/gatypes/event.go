package gatypes

// Event encapsulates information about a generic event
type Event struct { // t=event
	EventCategory string `url:"ec,omitempty" json:"event_category,omitempty"`
	EventAction   string `url:"ea,omitempty" json:"event_action,omitempty"`
	EventLabel    string `url:"el,omitempty" json:"event_label,omitempty"`
	EventValue    int64  `url:"ev,omitempty" json:"event_value,omitempty"`
}
