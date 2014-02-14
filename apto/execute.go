package apto

import (
	"os"
	"os/exec"
	"strings"
)

func Execute(command *Command) error {
	c := strings.Split(command.String(), " ")
	cmd := exec.Command(c[0], c[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
