package gatypes

// Timing encapsulates data available on a timing event
type Timing struct { // t=timing
	UserTimingCategory     string `url:"utc,omitempty" json:"user_timing_category,omitempty"`
	UserTimingVariableName string `url:"utv,omitempty" json:"user_timing_variable_name,omitempty"`
	UserTimingTime         int64  `url:"utt,omitempty" json:"user_timing_time,omitempty"`
	UserTimingLabel        string `url:"utl,omitempty" json:"user_timing_label,omitempty"`
	PageLoadTime           int64  `url:"plt,omitempty" json:"page_load_time,omitempty"`
	DNSTime                int64  `url:"dns,omitempty" json:"dns_time,omitempty"`
	PageDownloadTime       int64  `url:"pdt,omitempty" json:"page_download_time,omitempty"`
	RedirectResponseTime   int64  `url:"rrt,omitempty" json:"redirect_response_time,omitempty"`
	TCPConnectTime         int64  `url:"tcp,omitempty" json:"tcp_connect_time,omitempty"`
	ServerResponseTime     int64  `url:"srt,omitempty" json:"server_response_time,omitempty"`
	DOMInteractiveTime     int64  `url:"dit,omitempty" json:"dom_interactive_time,omitempty"`
	ContentLoadTime        int64  `url:"clt,omitempty" json:"content_load_time,omitempty"`
}
