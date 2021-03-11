package gatypes

// Exception encapsulates information about an exception event
type Exception struct { // t=exception
	ExceptionDescription string  `url:"exd,omitempty" json:"exception_description,omitempty"`
	IsExceptionFatal     Boolean `url:"exf,omitempty" json:"is_exception_fatal,omitempty"`
}
