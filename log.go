package log

import (
	"fmt"
	"io"
)

type Logger struct {
	Writer   io.Writer
	Prefixes []string
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	if nil == receiver.Writer {
		return
	}
	fmt.Fprintf(receiver.Writer, format+"\n", a...)
}

func (receiver Logger) Log(a ...interface{}) {
	s := fmt.Sprint(a...)
	receiver.Logf("%s", s)
}

func (receiver Logger) Begin() {
	receiver.Log("BEGIN")
}

func (receiver Logger) End() {
	receiver.Log("END")
}

func (receiver Logger) Prefix(newPrefix ...string) Logger {
	var newLogger = receiver
	newLogger.Prefixes = append(newLogger.Prefixes, newPrefix...)
	return newLogger
}
