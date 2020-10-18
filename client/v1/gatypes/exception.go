package gatypes

// Exception encapsulates information about an exception event
type Exception struct { // t=exception
	ExceptionDescription string  `url:"exd,omitempty"`
	IsExceptionFatal     Boolean `url:"exf,omitempty"`
}
