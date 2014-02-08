package apto

import (
	"fmt"
)

func Debug(msg string, i ...interface{}) {
	fmsg := fmt.Sprintf(msg, i...)
	fmt.Printf("--- DEBUG: %v\n", fmsg)
}
