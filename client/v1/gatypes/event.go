package gatypes

type Event struct { // t=event
	EventCategory string `url:"ec"`
	EventAction   string `url:"ea"`
	EventLabel    string `url:"el"`
	EventValue    int64  `url:"ev"`
}
