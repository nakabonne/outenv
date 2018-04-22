package main

import (
	"fmt"
	"log"
)

const withColor = "\033[31mdirenv: %s\033[0m"

func logError(msg string, a ...interface{}) {
	logF(withColor, msg, a...)
}

func logF(format, msg string, a ...interface{}) {
	defer log.SetFlags(log.Flags())
	defer log.SetPrefix(log.Prefix())
	log.SetFlags(0)
	log.SetPrefix("")

	msg = fmt.Sprintf(format+"\n", msg)
	log.Printf(msg, a...)
}
