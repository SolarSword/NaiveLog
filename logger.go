// the logger is to construct logs
//
package naivelog

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"
)

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

func New(opts ...Option) *logger {
	logger := &logger{opt: initOptions(opts...)}
	logger.entryPool = &sync.Pool{New: func() interface{} {
		return entry(logger)
	}}
	return logger
}

// this is the standard, default and global logger
var std = New()

func StdLogger() *logger {
	return std
}

// for the standard logger, the same belows
func SetOptions(opts ...Option) {
	std.SetOptions(opts...)
}

func (l *logger) SetOptions(opts ...Option) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}

/*
type Writer interface {
	Write(p []byte) (n int, err error)
}
*/

func Writer() io.Writer {
	return std
}

func (l *logger) Writer() io.Writer {
	return l
}

func (l *logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *logger) Write(data []byte) (int, error) {
	l.entry().write(l.opt.stdLevel, FmtEmptySeparate, *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func (l *logger) Debug(args ...interface{}) {
	l.entry().write(Debug, FmtEmptySeparate, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.entry().write(Info, FmtEmptySeparate, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.entry().write(Warn, FmtEmptySeparate, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.entry().write(Error, FmtEmptySeparate, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.entry().write(Panic, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func (l *logger) Fatal(args ...interface{}) {
	l.entry().write(Fatal, FmtEmptySeparate, args...)
	os.Exit(1)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry().write(Debug, format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.entry().write(Info, format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.entry().write(Warn, format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry().write(Error, format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.entry().write(Panic, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry().write(Fatal, format, args...)
	os.Exit(1)
}
