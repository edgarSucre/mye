package mye

type Err struct {
	t   ErrorType
	err error
}

func (e Err) Error() string {
	return e.err.Error()
}
