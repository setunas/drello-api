package apperr

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

func NewHTTPError(status int, message string, err error) error {
	switch e := err.(type) {
	case *AppError:
		return &HTTPError{
			status:     convertTagsToStatus(status, e.tags),
			message:    combineMessages(message, err),
			occurredAt: e.occurredAt,
		}
	case *HTTPError:
		return &HTTPError{
			status:     inheritStatus(status, e.status),
			message:    combineMessages(message, err),
			occurredAt: e.occurredAt,
		}
	default:
		return &HTTPError{
			status:     status,
			message:    combineMessages(message, err),
			occurredAt: newOccurredAt(),
		}
	}
}

func inheritStatus(newStatus int, oldStatus int) int {
	if oldStatus != 0 {
		return oldStatus
	} else {
		return newStatus
	}
}

func combineMessages(message string, err error) string {
	if message != "" && err != nil {
		return message + ": " + err.Error()
	} else if err != nil {
		return err.Error()
	} else {
		return message
	}
}

func newOccurredAt() string {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s:%d %s", file, line, f.Name())
}

var tagStatusMap = map[Tag]int{
	RecordNotFound:      404,
	FailedAuthorization: 403,
}

func convertTagsToStatus(newStatus int, tags []Tag) int {
	for _, tag := range tags {
		if oldStatus := tagStatusMap[tag]; oldStatus != 0 {
			return oldStatus
		}
	}
	return newStatus
}
