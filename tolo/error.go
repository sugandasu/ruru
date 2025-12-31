package tolo

import "errors"

type apperror struct {
	statuCode int
	message   string
	data      any
}

type Error interface {
	StatusCode() int
	Data() any
	error
}

func NewError(statuCode int, msg string, data any) Error {
	return &apperror{
		statuCode: statuCode,
		message:   msg,
		data:      data,
	}
}

func ErrorOr(cond bool, err1 Error, err2 Error) Error {
	if cond {
		return err1
	}

	return err2
}

func (e apperror) Error() string {
	return e.message
}

func (e apperror) StatusCode() int {
	return e.statuCode
}

func (e apperror) Data() any {
	return e.data
}

func ParseError(err error) Error {
	var apperr Error
	if !errors.As(err, &apperr) {
		return nil
	}
	return apperr
}
