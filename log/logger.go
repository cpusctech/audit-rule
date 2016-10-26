package log

import (
	"log"
	"runtime/debug"
	"os"
)

var loggerFlag = log.Lshortfile | log.Ldate | log.Ltime

type Level int

/**
Java-Like Logger
*/
const (
	TRACE = Level(0)
	DEBUG = Level(1)
	INFO  = Level(2)
	WARN  = Level(3)
	ERROR = Level(4)
)

type Logger struct {
	*log.Logger
	level Level
}

func (gl *Logger) SetLevel(l Level) {
	gl.level = l
}

func (gl *Logger) Trace(format string, args ...interface{}) {
	if gl.level <= TRACE {
		gl.Logger.Printf("[TRACE]"+format, args...)
	}
}

func (gl *Logger) Debug(format string, args ...interface{}) {
	if gl.level <= DEBUG {
		gl.Logger.Printf("[DEBUG]"+format, args...)
	}
}

func (gl *Logger) Info(format string, args ...interface{}) {
	if gl.level <= INFO {
		gl.Logger.Printf("[INFO]"+format, args...)
	}
}

func (gl *Logger) Warn(format string, args ...interface{}) {
	if gl.level <= WARN {
		gl.Logger.Printf("[WARN]"+format, args...)
	}
}

// NOTE: 只有ERROR的时候会把Stack Trace打出来
func (gl *Logger) Error(format string, args ...interface{}) {
	if gl.level <= ERROR {
		stack := debug.Stack()
		gl.Logger.Printf("[ERROR]"+format+"\nstack trace: "+string(stack), args...)
	}
}

func (gl *Logger) ErrorWithoutStack(format string, args ...interface{}) {
	if gl.level <= ERROR {
		gl.Logger.Printf("[ERROR]"+format, args...)
	}
}

var CustomLogger = &Logger{
	Logger: log.New(os.Stdout, "", loggerFlag),
	level:  DEBUG,
}