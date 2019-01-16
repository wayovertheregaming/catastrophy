// +build !debug

package catlog

import (
	"log"
	"os"
)

func init() {
	myLog = log.New(
		os.Stdout,
		"Running in Prod mode",
		log.LstdFlags,
	)
}

// Debug will write a debug line to the log
func Debug(_ ...interface{}) {}

// Debugf will format and write a debug line to the log
func Debugf(_ string, _ ...interface{}) {}

// Info will write an info line to the log
func Info(_ ...interface{}) {}

// Infof will write and format an info line to the log
func Infof(_ string, _ ...interface{}) {}
