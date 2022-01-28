package myerr

import (
	"fmt"
	"runtime"
)

type HTTPError struct {
	status     int
	message    string
	occurredAt string
}

func (e *HTTPError) Error() string {
	return e.message
}

func (e *HTTPError) Status() int {
	return e.status
}

func (e *HTTPError) OccurredAt() string {
	return e.occurredAt
}

func (e *HTTPError) IsClientError() bool {
	return 400 <= e.Status() && e.Status() < 500
}

func NewHTTPError(status int, detail string, err error) error {
	httpError, ok := err.(*HTTPError)
	if ok {
		return &HTTPError{
			status:     inheritStatus(status, httpError.status),
			message:    newMessage(detail, err),
			occurredAt: httpError.occurredAt,
		}
	}

	return &HTTPError{
		status:     status,
		message:    newMessage(detail, err),
		occurredAt: newOccurredAt(),
	}
}

func inheritStatus(newStatus int, oldStatus int) int {
	if oldStatus != 0 {
		return oldStatus
	} else {
		return newStatus
	}
}

func newOccurredAt() string {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	place := fmt.Sprintf("%s:%d %s\n", file, line, f.Name())
	return place
}

func newMessage(message string, err error) string {
	if message != "" && err != nil {
		return message + ": " + err.Error()
	} else if err != nil {
		return err.Error()
	} else {
		return message
	}
}
