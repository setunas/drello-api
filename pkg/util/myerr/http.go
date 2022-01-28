package myerr

import (
	"fmt"
	"runtime"
)

type HTTPError struct {
	status int
	detail string
	place  string
}

func (e *HTTPError) Error() string {
	return e.detail
}

func (e *HTTPError) Status() int {
	return e.status
}

func (e *HTTPError) Place() string {
	return e.place
}

func (e *HTTPError) IsClientError() bool {
	return 400 <= e.Status() && e.Status() < 500
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
			status: newStatus(status, httpError.status),
			detail: newDetail(detail, err),
			place:  httpError.place,
		}
	}

	return &HTTPError{
		status: status,
		detail: newDetail(detail, err),
		place:  newPlace(),
	}
}
