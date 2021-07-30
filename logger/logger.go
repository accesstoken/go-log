package logger

import (
	"fmt"
	"io"
)

type Logger struct {
	Writer io.Writer
}

func (receiver Logger) Log(a ...interface{}) {
	fmt.Fprintln(receiver.Writer,  a...)
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	fmt.Fprintf(receiver.Writer,  format, a...)
}

func (receiver Logger) Begin(a ...interface{}) {
	receiver.Log("BEGIN")
}

func (receiver Logger) End(a ...interface{}) {
	receiver.Log("END")
}
