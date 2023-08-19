package naivelog

import (
	"io"
	"os"
)

const (
	FmtEmptySeparate = ""
)

type Level uint8 // log level

const (
	Debug Level = iota
	Info
	Warn
	Error
	Panic
	Fatal
)

var LevelNameMap = map[Level]string{
	Debug: "DEBUG",
	Info:  "INFO",
	Warn:  "WARN",
	Error: "ERROR",
	Panic: "PANIC",
	Fatal: "FATAL",
}

type options struct {
	output        io.Writer
	level         Level // log level
	stdLevel      Level // standard log level
	formatter     Formatter
	disableCaller bool
}

type Option func(*options)

func initOptions(opts ...Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	return o
}

func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithStdLevel(level Level) Option {
	return func(o *options) {
		o.stdLevel = level
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

func WithDisableCaller(caller bool) Option {
	return func(o *options) {
		o.disableCaller = caller
	}
}
