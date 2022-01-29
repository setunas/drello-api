package log

import (
	"drello-api/pkg/util/color"
	"fmt"
	"log"
	"os"
)

type Log struct {
	level   Level
	message string
	fields  map[string]string
}

func (l *Log) Add(key, value string) *Log {
	l.fields[key] = value
	return l
}

func (l *Log) Write() {
	var output string
	switch l.level {
	case fatal:
		output = color.Red + "[!!!FATAL!!!] " + color.Reset
	case err:
		output = color.Red + "[ERROR] " + color.Reset
	case warn:
		output = color.Yellow + "[WARN] " + color.Reset
	default:
		output = color.Green + "[INFO] " + color.Reset
	}

	output += l.message

	if l.fields != nil {
		output += " { "

		prefix := ""
		for k, v := range l.fields {
			key := color.Green + k + color.Reset
			value := color.Yellow + v + color.Reset
			output += prefix + key + ": " + value
			prefix = ", "
		}

		output += " }"
	}

	log.Println(output)

	if l.level == fatal {
		os.Exit(1)
	}
}

type Level int

const (
	info Level = iota
	warn
	err
	fatal
)

func Info(values ...interface{}) *Log {
	return &Log{
		level:   info,
		message: fmt.Sprint(values...),
		fields:  map[string]string{},
	}
}

func Warn(values ...interface{}) *Log {
	return &Log{
		level:   warn,
		message: fmt.Sprint(values...),
		fields:  make(map[string]string),
	}
}

func Err(values ...interface{}) *Log {
	return &Log{
		level:   err,
		message: fmt.Sprint(values...),
		fields:  make(map[string]string),
	}
}

func Fatal(values ...interface{}) *Log {
	return &Log{
		level:   fatal,
		message: fmt.Sprint(values...),
		fields:  make(map[string]string),
	}
}
