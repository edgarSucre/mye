package mye

import "errors"

type ErrorType int

const (
	Cancelation ErrorType = iota
	Forbiden
	Internal
	NotFound
	Timeout
	Unauthorized
	Validation
)

var (
	alertable = []ErrorType{Forbiden, Internal, Timeout}
	loggeable = []ErrorType{Internal, Timeout}
)

func (et ErrorType) New(msg string) error {
	return Err{t: et, err: errors.New(msg)}
}

func (et ErrorType) isAlertable() bool {
	return et.isIn(alertable)
}

func (et ErrorType) isLoggeable() bool {
	return et.isIn(loggeable)
}

func (et ErrorType) isIn(types []ErrorType) bool {
	for _, v := range types {
		if et == v {
			return true
		}
	}

	return false
}
