package apto

import (
	"fmt"
)

// Debug prints with a custom format
func Debug(msg string, i ...interface{}) {
	fmsg := fmt.Sprintf(msg, i...)
	fmt.Printf("--- DEBUG: %v\n", fmsg)
}

// HandleFlag looks for a flag within args,
// returns args without given flag and bool if flag was in args
func HandleFlag(args []string, flag string) ([]string, bool) {
	var exists bool
	var index int
	var arg string

	for index, arg = range args {
		if arg == flag {
			exists = true
			break
		}
	}

	if exists {
		args = append(args[:index], args[index+1:]...)
	}

	return args, exists
}
