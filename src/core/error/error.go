package error

import (
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
)

type appErr struct {
	error
	msg    string
	trace  error
	status int
}

type appError interface {
	TraceString() string
	Error() string
	Status() int
	AppendTrace(err error)
}

func (e *appErr) TraceString() string {
	if e.trace == nil {
		return ""
	}
	return e.trace.Error()
}

func (e *appErr) Error() string {
	return e.msg
}

func (e *appErr) Status() string {
	return e.msg
}

func (e *appErr) AppendTrace() string {
	return e.msg
}

type BadRequestErr struct {
	*appErr
}

type NotFoundErr struct {
	*appErr
}

type InternalServerErr struct {
	*appErr
}

func BadRequestWrapf(format string, args ...interface{}) error {
	err := &appErr{
		status: http.StatusBadRequest,
	}

	if len(args) == 0 {
		err.msg = format
		err.trace = errors.Errorf("%+v", errors.New(format))
		return &BadRequestErr{err}
	}

	err.msg = fmt.Sprintf(format, args...)
	err.trace = errors.Errorf("%+v", errors.Errorf(format, args...))
	return &BadRequestErr{err}
}

func NotFoundWrapf(format string, args ...interface{}) error {
	err := &appErr{
		status: http.StatusNotFound,
	}

	if len(args) == 0 {
		err.msg = format
		err.trace = errors.Errorf("%+v", errors.New(format))
		return &NotFoundErr{err}
	}

	err.msg = fmt.Sprintf(format, args...)
	err.trace = errors.Errorf("%+v", errors.Errorf(format, args...))
	return &NotFoundErr{err}
}

func InternalServerWrapf(format string, args ...interface{}) error {
	err := &appErr{
		status: http.StatusInternalServerError,
	}

	if len(args) == 0 {
		err.msg = format
		err.trace = errors.Errorf("%+v", errors.New(format))
		return &InternalServerErr{err}
	}

	err.msg = fmt.Sprintf(format, args...)
	err.trace = errors.Errorf("%+v", errors.Errorf(format, args...))
	return &InternalServerErr{err}
}
