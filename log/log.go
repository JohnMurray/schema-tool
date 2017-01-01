package log

/**
 * Define some simple logging infastrucutre that can be used from anywhere else in
 * the program. To use simply:
 *
 *    log.LEVEL.Println("my log message")
 *
 * Where LEVEL can be any one of:
 *    - Trace
 *    - Info
 *    - Warning
 *    - Error
 */

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// trace logger instance
var Trace *log.Logger

// info logger instance
var Info *log.Logger

// warning logger instance
var Warn *log.Logger

// error logging instance
var Error *log.Logger

// simple object to ensure that we only run the InitLoggers function once
// even if called multiple times
var once sync.Once

// Initialize loggers (varies based on whether or not we're "trace" logging)
func InitLoggers(traceLoggingEnabled bool) {
	shouldInit := false
	once.Do(func() { shouldInit = true })
	if !shouldInit {
		return
	}

	infoHandle := os.Stderr
	warningHandle := os.Stderr
	errorHandle := os.Stderr

	var traceHandle io.Writer
	var logFormat int

	if traceLoggingEnabled {
		traceHandle = os.Stderr
		logFormat = log.Ldate | log.Ltime | log.Lshortfile
	} else {
		traceHandle = ioutil.Discard
		logFormat = log.Ldate | log.Ltime
	}

	Trace = log.New(traceHandle, "TRACE: ", logFormat)
	Info = log.New(infoHandle, "INFO: ", logFormat)
	Warn = log.New(warningHandle, "WARNING: ", logFormat)
	Error = log.New(errorHandle, "ERROR: ", logFormat)
}

// Initialize loggers that discard results. Useful to quiet logs during testing
func InitNopLoggers() {
	nopHandle := ioutil.Discard

	Trace = log.New(nopHandle, "", 0)
	Info = log.New(nopHandle, "", 0)
	Warn = log.New(nopHandle, "", 0)
	Error = log.New(nopHandle, "", 0)
}
