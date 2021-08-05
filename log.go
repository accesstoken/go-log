package log

import (
	"fmt"
	"io"
	"strings"
)

type Logger struct {
	Writer   io.Writer
	prefixes []string
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	if nil == receiver.Writer {
		return
	}
	if receiver.prefixes != nil {
		var prefixString string = strings.Join(receiver.prefixes[:], ": ") + ":"
		a = append([]interface{}{prefixString}, a...)
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
	newLogger.prefixes = append(newLogger.prefixes, newPrefix...)
	return newLogger
}
