package mye

import (
	"errors"
	"fmt"
)

type Err struct {
	T   ErrorType
	err error
}

func (e Err) Error() string {
	return e.err.Error()
}

func IsLoggeable(err error) bool {
	if local, ok := err.(Err); ok {
		return local.T.isLoggeable()
	}

	return true
}

func IsAlertable(err error) bool {
	if local, ok := err.(Err); ok {
		return local.T.isAlertable()
	}

	return true
}

func Wrap(err error, msg string) error {
	return fmt.Errorf("%w %s", err, msg)
}

func Wrapf(err error, format string, obs ...any) error {
	msg := fmt.Sprintf(format, obs...)
	return Wrap(err, msg)
}

func UnWrap(err error) error {
	if unwraped := errors.Unwrap(err); unwraped != nil {
		UnWrap(unwraped)
	}

	return err
}
