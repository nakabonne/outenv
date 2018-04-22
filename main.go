package main

import (
	"os"
)

// Configured at compile time
var bashPath string

func main() {
	var env = GetEnv()
	var args = os.Args

	err := CommandsDispatch(env, args)
	if err != nil {
		log_error("error %v", err)
		os.Exit(1)
	}
}
