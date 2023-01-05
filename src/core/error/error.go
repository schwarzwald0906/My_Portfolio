package error

import (
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
)

type appErr struct {
	error
	msg      string
	trace    error
	statusCd int
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

func BadRequestWrapf(format string, any interface{}) error {
	err := &appErr{
		statusCd: http.StatusBadRequest,
	}

	if any == nil {
		err.msg = format
		err.trace = errors.Errorf("%+v", errors.New(format))
		return &BadRequestErr{err}
	}

	err.msg = fmt.Sprintf(format, any)
	err.trace = errors.Errorf("%+v", errors.Errorf(format, any))
	return &BadRequestErr{err}
}

func NotFoundWrapf(format string, any interface{}) error {
	err := &appErr{
		statusCd: http.StatusNotFound,
	}

	if any == nil {
		err.msg = format
		err.trace = errors.Errorf("%+v", errors.New(format))
		return &NotFoundErr{err}
	}

	err.msg = fmt.Sprintf(format, any)
	err.trace = errors.Errorf("%+v", errors.Errorf(format, any))
	return &NotFoundErr{err}
}

func InternalServerWrapf(format string, any interface{}) error {
	err := &appErr{
		statusCd: http.StatusInternalServerError,
	}

	if any == nil {
		err.msg = format
		err.trace = errors.Errorf("%+v", errors.New(format))
		return &InternalServerErr{err}
	}

	err.msg = fmt.Sprintf(format, any)
	err.trace = errors.Errorf("%+v", errors.Errorf(format, any))
	return &InternalServerErr{err}
}
