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

func (e *HTTPError) Error() string {
	return e.Detail
}

func newStatus(status int, oldStatus int) int {
	if oldStatus != 0 {
		return oldStatus
	} else {
		return status
	}
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
			Status: newStatus(status, httpError.Status),
			Detail: newDetail(detail, err),
			Place:  httpError.Place,
		}
	}

	return &HTTPError{
		Status: status,
		Detail: newDetail(detail, err),
		Place:  newPlace(),
	}
}
