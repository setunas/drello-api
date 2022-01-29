package apperr

import (
	"fmt"
	"runtime"
)

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
	_, filename, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", filename, line)
}
