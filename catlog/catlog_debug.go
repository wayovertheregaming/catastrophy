// +build debug

package catlog

import (
	"fmt"
	"log"
	"os"
)

func init() {
	myLog = log.New(
		os.Stdout,
		"Running in debug mode",
		log.Ldate|log.Ltime|log.Llongfile,
	)
}

// Debug will write a debug line to the log
func Debug(v ...interface{}) {
	myLog.Printf("[DEBUG] %v", v...)
}

// Debugf will format and write a debug line to the log
func Debugf(format string, v ...interface{}) {
	formatted := fmt.Sprintf(format, v...)
	myLog.Printf("[DEBUG] %s", formatted)
}

// Info will write an info line to the log
func Info(v ...interface{}) {
	myLog.Printf("[INFO] %v", v...)
}

// Infof will write and format an info line to the log
func Infof(format string, v ...interface{}) {
	formatted := fmt.Sprintf(format, v...)
	myLog.Printf("[Info] %s", formatted)
}
