package gatypes

// Event encapsulates information about a generic event
type Event struct { // t=event
	EventCategory string `url:"ec,omitempty"`
	EventAction   string `url:"ea,omitempty"`
	EventLabel    string `url:"el,omitempty"`
	EventValue    int64  `url:"ev,omitempty"`
}
