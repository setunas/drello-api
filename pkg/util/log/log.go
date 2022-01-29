package log

import (
	"drello-api/pkg/util/color"
	"log"
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
	case err:
		output = color.Red + "[ERROR] " + color.Reset
	case warn:
		output = color.Yellow + "[WARN] " + color.Reset
	case info:
		output = "[INFO] "
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
}

type Level int

const (
	info Level = iota
	warn
	err
)

func Info(message string) *Log {
	return &Log{
		level:   info,
		message: message,
		fields:  map[string]string{},
	}
}

func Warn(message string) *Log {
	return &Log{
		level:   warn,
		message: message,
		fields:  make(map[string]string),
	}
}

func Err(message string) *Log {
	return &Log{
		level:   err,
		message: message,
		fields:  make(map[string]string),
	}
}
