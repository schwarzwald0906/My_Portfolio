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

func BadRequestWrapf(format string, args ...interface{}) error {
	err := &appErr{
		statusCd: http.StatusBadRequest,
	}

	if len(args) == 0 {
		err.msg = format
		//提示いただいた例だと、errors.New()をさらにErrorf（）してましたが、
		//Errorf（)を挟む理由がよくわからなかったので、直でやりました。
		err.trace = errors.New(format)
		return &BadRequestErr{err}
	}

	err.msg = fmt.Sprintf(format, args...)
	//ここも、Errorf（)にErrorf（)を挟む理由がよくわからなかったので、直でやりました。
	err.trace = errors.Errorf(format, args...)
	return &BadRequestErr{err}
}

func NotFoundWrapf(format string, args ...interface{}) error {
	err := &appErr{
		statusCd: http.StatusNotFound,
	}

	if len(args) == 0 {
		err.msg = format
		err.trace = errors.New(format)
		return &NotFoundErr{err}
	}

	err.msg = fmt.Sprintf(format, args...)
	err.trace = errors.Errorf(format, args...)
	return &NotFoundErr{err}
}

func InternalServerWrapf(format string, args ...interface{}) error {
	err := &appErr{
		statusCd: http.StatusInternalServerError,
	}

	if len(args) == 0 {
		err.msg = format
		err.trace = errors.New(format)
		return &InternalServerErr{err}
	}

	err.msg = fmt.Sprintf(format, args...)
	err.trace = errors.Errorf(format, args...)
	return &InternalServerErr{err}
}
