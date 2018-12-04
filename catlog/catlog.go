package catlog

import (
	"fmt"
	"log"
)

var (
	myLog *log.Logger
)

// Fatal will write a fatal line to the log, then exit the application
func Fatal(v ...interface{}) {
	myLog.Fatalf("[FATAL ERROR] %v", v...)
}

// Fatalf will write and format a fatal line to the log, then exit the
// application
func Fatalf(format string, v ...interface{}) {
	formatted := fmt.Sprintf(format, v...)
	myLog.Fatalf("[FATAL ERROR] %s", formatted)
}
