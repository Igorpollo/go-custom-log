package logo

import (
    "time"
    "fmt"
	"os"
)

const timeFormat = "%s %s %s\n"
const simpleFormat = "%s %s\n"

const (
    _info = iota
	_warning
	_success
    _error
    _debug
)

var codeMsgs = [5]string {
    " \x1b[37;1m[INFO]\x1b[0m",
    " \x1b[33;1m[WARNING]\x1b[0m",
	" \x1b[32;1m[SUCCESS]\x1b[0m",
	" \x1b[31;1m[ERROR]\x1b[0m",
	" \x1b[34;1;3m[DEBUG]\x1b[0m",
}

var debug bool
var showTime bool

var logLevel = _info

func init() {
	verifyDebugMode()
}

func verifyDebugMode() {
	if os.Getenv("DEBUG") == "true" {
		debug = true
	}
}

// DebugMode set if log.Debug will appear
func DebugMode(flag bool) {
	debug = flag
}

//ShowTime define if time will appear
func ShowTime(flag bool) {
	showTime = flag
}

func Success(format string, v ...interface{}) {
    logger(_success, fmt.Sprintf(format, v ...))
}

func Debug(format string, v ...interface{}) {
	if debug == false {
		return
	}
    logger(_debug, fmt.Sprintf(format, v ...))
}

func Info(format string, v ...interface{}) {
    logger(_info, fmt.Sprintf(format, v ...))
}

func Warning(format string, v ...interface{}) {
    logger(_warning, fmt.Sprintf(format, v ...))
}

func Error(format string, v ...interface{}) {
    logger(_error, fmt.Sprintf(format, v ...))
}

func logger(code int, message string) {
	timeSting := fmt.Sprintf("%s", time.Now())[:19]
	if showTime == true {
		fmt.Printf(timeFormat, codeMsgs[code], timeSting, message)
	} else {
		fmt.Printf(simpleFormat, codeMsgs[code], message)
	}
}