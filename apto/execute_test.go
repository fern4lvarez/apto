package apto

import (
	"testing"
)

func TestExecute(t *testing.T) {
	spec := "Should execute not sudo bash command and return output"
	command := &Command{Sudo: false,
		Tool:    "apt-get",
		Cmd:     "help",
		Pkgs:    []string{},
		Options: []string{}}

	if err := Execute(command); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}

func TestExecuteError(t *testing.T) {
	spec := "Should return error when executing a wrong bash command"
	command := &Command{Sudo: false,
		Tool:    "notexists",
		Cmd:     "barbarbar",
		Pkgs:    []string{},
		Options: []string{}}

	if err := Execute(command); err == nil {
		t.Errorf(msg, spec, err, nil)
	}
}
