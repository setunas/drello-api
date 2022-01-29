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
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s:%d %s", file, line, f.Name())
}
