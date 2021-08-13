package log

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const PackagePath = "github.com/accesstoken/go-log.Logger."

type Logger struct {
	Writer    io.Writer
	prefixes  []string
	startTime time.Time
	level     uint8
}

func funcName() string {

	pc := make([]uintptr, 10)
	n := runtime.Callers(2, pc)
	if n == 0 {
		return "?()"
	}
	var fnName string
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		fnName = frame.Function
		if !more || !strings.HasPrefix(fnName, PackagePath) {
			fnName = strings.TrimLeft(filepath.Ext(fnName), ".") + "()"
			break
		}
	}
	return fnName
}

func (receiver Logger) writef(format string, a ...interface{}) {

	writer := receiver.Writer

	if nil == writer {
		return
	}

	s := fmt.Sprintf(format, a...)

	if receiver.prefixes != nil {
		prefixString := strings.Join(receiver.prefixes[:], ": ") + ": "
		s = prefixString + s
	}

	io.WriteString(writer, funcName()+" -> "+s+"\n")

}

func (receiver Logger) write(a ...interface{}) {
	s := fmt.Sprint(a...)
	receiver.writef("%s", s)
}

func (receiver Logger) Alert(a ...interface{}) {
	if receiver.level >= 1 {
		receiver.write(a...)
	}
}

func (receiver Logger) Alertf(format string, a ...interface{}) {
	if receiver.level >= 1 {
		receiver.writef(format, a...)
	}
}

func (receiver Logger) Error(a ...interface{}) {
	if receiver.level >= 1 {
		receiver.write(a...)
	}
}

func (receiver Logger) Errorf(format string, a ...interface{}) {
	if receiver.level >= 1 {
		receiver.writef(format, a...)
	}
}

func (receiver Logger) Warn(a ...interface{}) {
	if receiver.level >= 2 {
		receiver.write(a...)
	}
}

func (receiver Logger) Warnf(format string, a ...interface{}) {
	if receiver.level >= 2 {
		receiver.writef(format, a...)
	}
}

func (receiver Logger) Highlight(a ...interface{}) {
	if receiver.level >= 3 {
		receiver.write(a...)
	}
}

func (receiver Logger) Highlightf(format string, a ...interface{}) {
	if receiver.level >= 3 {
		receiver.writef(format, a...)
	}
}

func (receiver Logger) Inform(a ...interface{}) {
	if receiver.level >= 4 {
		receiver.write(a...)
	}
}

func (receiver Logger) Informf(format string, a ...interface{}) {
	if receiver.level >= 4 {
		receiver.writef(format, a...)
	}
}

func (receiver Logger) Log(a ...interface{}) {
	if receiver.level >= 5 {
		receiver.write(a...)
	}
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	if receiver.level >= 5 {
		receiver.writef(format, a...)
	}
}

func (receiver Logger) Trace(a ...interface{}) {
	if receiver.level >= 6 {
		receiver.write(a...)
	}
}

func (receiver Logger) Tracef(format string, a ...interface{}) {
	if receiver.level >= 6 {
		receiver.writef(format, a...)
	}
}

func (receiver Logger) Level(level uint8) Logger {
	var newLogger = receiver
	newLogger.level = level
	return newLogger
}

func (receiver Logger) Begin() Logger {
	var newLogger = receiver
	newLogger.startTime = time.Now()
	newLogger.write("BEGIN")
	return newLogger
}

func (receiver Logger) End() {
	elapsed := time.Since(receiver.startTime)
	receiver.writef("END δt=%v", elapsed)
}

func (receiver Logger) Prefix(newPrefix ...string) Logger {
	var newLogger = receiver
	newLogger.prefixes = append(newLogger.prefixes, newPrefix...)
	return newLogger
}
