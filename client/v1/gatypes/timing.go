package gatypes

type Timing struct { // t=timing
	UserTimingCategory     string `url:"utc,omitempty"`
	UserTimingVariableName string `url:"utv,omitempty"`
	UserTimingTime         int64  `url:"utt,omitempty"`
	UserTimingLabel        string `url:"utl,omitempty"`

	// less often used

	PageLoadTime         int64 `url:"plt,omitempty"`
	DNSTime              int64 `url:"dns,omitempty"`
	PageDownloadTime     int64 `url:"pdt,omitempty"`
	RedirectResponseTime int64 `url:"rrt,omitempty"`
	TCPConnectTime       int64 `url:"tcp,omitempty"`
	ServerResponseTime   int64 `url:"srt,omitempty"`
	DOMInteractiveTime   int64 `url:"dit,omitempty"`
	ContentLoadTime      int64 `url:"clt,omitempty"`
}
