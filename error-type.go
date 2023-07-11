package mye

type ErrorType int

const (
	Cancelation = iota
	Forbiden
	Internal
	NotFound
	Timeout
	Unauthorized
	Validation
)

var (
	reportable = []ErrorType{Forbiden, Internal, Timeout}
	loggeable  = []ErrorType{Internal, Timeout}
)

func (et ErrorType) New(err error) error {
	return Err{t: et, err: err}
}

func (et ErrorType) IsReportable() bool {
	return et.isIn(reportable)
}

func (et ErrorType) IsLoggeable() bool {
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
