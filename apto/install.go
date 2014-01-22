package apto

import (
	"errors"
)

func Install(args []string) error {
	command, err := installCommand(args)
	if err != nil {
		return err
	}

	return Execute(command)
}

func installCommand(args []string) (*Command, error) {
	sudo := true
	tool := "apt-get"
	cmd := args[1] // install
	pkgs := args[2:]
	if len(pkgs) == 0 {
		return nil, errors.New("You must provide packages to install.")
	}
	options := []string{"-y"}

	return NewCommand(sudo, tool, cmd, pkgs, options)
}
