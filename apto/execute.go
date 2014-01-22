package apto

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func Execute(command *Command) error {
	c := strings.Split(command.String(), " ")
	cmd := exec.Command(c[0], c[1:]...)
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		return err
	}

	go io.Copy(os.Stdout, stdout)

	return cmd.Wait()
}
