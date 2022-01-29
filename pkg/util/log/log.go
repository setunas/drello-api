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

func (l *Log) Add(key, value interface{}) *Log {
	l.fields = append(l.fields, &Field{key: fmt.Sprint(key), value: fmt.Sprint(value)})
	return l
}

func (l *Log) Write() {
	var output string
	switch l.level {
	case fatal:
		output = color.Red + "[!!!FATAL!!!] " + l.message + color.Reset
	case err:
		output = color.Red + "[ERROR] " + l.message + color.Reset
	case warn:
		output = color.Yellow + "[WARN] " + color.Reset + l.message
	default:
		output = color.Cyan + "[INFO] " + color.Reset + l.message
	}

	if len(l.fields) != 0 {
		output += " { "

		prefix := ""
		for _, v := range l.fields {
			key := color.Green + v.key + color.Reset
			value := color.Purple + v.value + color.Reset
			output += prefix + key + ": " + value
			prefix = ", "
		}

		output += " }"
	}

	log.Println(output)
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

func Infof(format string, values ...interface{}) *Log {
	return &Log{
		level:   info,
		message: fmt.Sprintf(format, values...),
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

func Warnf(format string, values ...interface{}) *Log {
	return &Log{
		level:   warn,
		message: fmt.Sprintf(format, values...),
		fields:  nil,
	}
}

func Error(values ...interface{}) *Log {
	return &Log{
		level:   err,
		message: fmt.Sprint(values...),
		fields:  nil,
	}
}

func Errorf(format string, values ...interface{}) *Log {
	return &Log{
		level:   err,
		message: fmt.Sprintf(format, values...),
		fields:  nil,
	}
}

func Fatal(values ...interface{}) {
	l := &Log{
		level:   fatal,
		message: fmt.Sprint(values...),
		fields:  nil,
	}

	l.Write()
	os.Exit(1)
}

func Fatalf(format string, values ...interface{}) {
	l := &Log{
		level:   fatal,
		message: fmt.Sprintf(format, values...),
		fields:  nil,
	}

	l.Write()
	os.Exit(1)
}
