package log

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	Writer   io.Writer
	prefixes []string
	startTime time.Time
}

func (receiver Logger) Logf(format string, a ...interface{}) {

	writer := receiver.Writer

	if nil == writer {
		return
	}

	s := fmt.Sprintf(format, a...)

	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		fnName = strings.TrimLeft(filepath.Ext(fn.Name()), ".") + "()"
	}

	s = fnName + " " + s

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

func (receiver Logger) Begin() Logger{
	var newLogger = receiver
	newLogger.startTime = time.Now()
	newLogger.Log("BEGIN")
	return newLogger
}

func (receiver Logger) End() {
	elapsed := time.Since(receiver.startTime)
	receiver.Log("END δt=", elapsed)
}

func (receiver Logger) Prefix(newPrefix ...string) Logger {
	var newLogger = receiver
	newLogger.prefixes = append(newLogger.prefixes, newPrefix...)
	return newLogger
}
