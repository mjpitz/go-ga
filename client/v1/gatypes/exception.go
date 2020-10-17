package gatypes

type Exception struct { // t=exception
	ExceptionDescription string  `url:"edx"`
	IsExceptionFatal     Boolean `url:"exf"`
}
