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

	writer := receiver.Writer

	if nil == writer {
		return
	}

	s := fmt.Sprintf(format, a...)

	if receiver.prefixes != nil {
		prefixString := strings.Join(receiver.prefixes[:], ": ") + ": "
		s = prefixString + s
	}
	_, err := io.WriteString(writer, s + "\n")

	if err != nil {
		return
	}
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
