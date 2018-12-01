package catlog

import "log"

var (
	myLog *log.Logger
)

// Fatal will write a fatal line to the log, then exit the application
func Fatal(v ...interface{}) {
	myLog.Fatal(v...)
}

// Fatalf will write and format a fatal line to the log, then exit the
// application
func Fatalf(format string, v ...interface{}) {
	myLog.Fatalf(format, v...)
}
