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
	fields  []*Field
}

type Field struct {
	key   string
	value string
}

func (l *Log) Add(key, value string) *Log {
	l.fields = append(l.fields, &Field{key: key, value: value})
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

	if len(l.fields) != 0 {
		output += " { "

		prefix := ""
		for _, v := range l.fields {
			key := color.Green + v.key + color.Reset
			value := color.Yellow + v.value + color.Reset
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
		fields:  nil,
	}
}

func Warn(values ...interface{}) *Log {
	return &Log{
		level:   warn,
		message: fmt.Sprint(values...),
		fields:  nil,
	}
}

func Err(values ...interface{}) *Log {
	return &Log{
		level:   err,
		message: fmt.Sprint(values...),
		fields:  nil,
	}
}

func Fatal(values ...interface{}) *Log {
	return &Log{
		level:   fatal,
		message: fmt.Sprint(values...),
		fields:  nil,
	}
}
