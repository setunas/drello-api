package myerr

import (
	"fmt"
	"runtime"
)

type HTTPError struct {
	Status int    `json:"-"`
	Detail string `json:"detail"`
	Cause  error  `json:"-"`
	Place  string
}

type HTTPErrorIF interface {
	Error() string
	GetStatus() int
	GetPlace() string
	ResponseBody() string
}

func (e *HTTPError) Error() string {
	return e.Detail
}

func (e *HTTPError) GetStatus(status int, err error) int {
	_, ok := err.(HTTPErrorIF)
	if !ok {
		return status
	}

	if e.Status != 0 {
		return e.Status
	}

	return status
}

func (e *HTTPError) GetPlace() string {
	return e.Place
}

func newPlace() string {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	place := fmt.Sprintf("%s:%d %s\n", file, line, f.Name())
	return place
}

func newDetail(detail string, err error) string {
	if detail != "" && err != nil {
		return detail + ": " + err.Error()
	} else if err != nil {
		return err.Error()
	} else {
		return detail
	}
}

func NewHTTPError(status int, detail string, err error) error {
	httpError, ok := err.(*HTTPError)
	if ok {
		return &HTTPError{
			Status: httpError.GetStatus(status, err),
			Detail: newDetail(detail, err) + ": " + httpError.Error(),
			Place:  httpError.GetPlace(),
		}
	}

	return &HTTPError{
		Status: status,
		Detail: newDetail(detail, err),
		Place:  newPlace(),
	}
}
