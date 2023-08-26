package core

import "fmt"

type TempletError interface {
	Error() error
	Cause() string
	String() string
}

type baseError struct {
	err   error
	cause string
}

func (e baseError) Error() error {
	return e.err
}

func (e baseError) Cause() string {
	return e.cause
}

func (e baseError) String() string {
	return fmt.Sprintf("[%s] %s", e.cause, e.err.Error())
}

func NewError(err error, cause string) TempletError {
	return baseError{
		err:   err,
		cause: cause,
	}
}
