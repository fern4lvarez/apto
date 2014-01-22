package apto

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

var (
	msg = "%v. Expects %v, returns %v"
)

func TestNewCommand(t *testing.T) {
	spec := "Should return a new Command"
	expectedCommand := Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"git-essentials"},
		Options: []string{},
	}

	if command, err := NewCommand(true,
		"apt-get",
		"install",
		[]string{"git-essentials"},
		[]string{}); err != nil {
		t.Errorf(msg, spec, nil, err)
	} else if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}
}

func TestNewCommandError(t *testing.T) {
	spec := "Should return error when Tool is empty"
	expectedErr := errors.New("Tool is empty.")
	if _, err := NewCommand(true,
		"",
		"install",
		[]string{"git-essentials"},
		[]string{}); !reflect.DeepEqual(expectedErr, err) {
		t.Errorf(msg, spec, expectedErr, err)
	}

	spec = "Should return error when Cmd is empty"
	expectedErr = errors.New("Cmd is empty.")
	if _, err := NewCommand(true,
		"apt-get",
		"",
		[]string{"git-essentials"},
		[]string{}); !reflect.DeepEqual(expectedErr, err) {
		t.Errorf(msg, spec, expectedErr, err)
	}

}

func TestCommandString(t *testing.T) {
	spec := "Should print Command with Sudo into a valid bash sudo command"
	command := &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"git-essentials", "vim"},
		Options: []string{"-y", "-d"},
	}
	expectedString := "sudo apt-get install git-essentials vim -y -d"

	if s := command.String(); s != expectedString {
		t.Errorf(msg, spec, expectedString, s)
	}

	spec = "Should print Command without sudo into a valid bash sudo command"
	command.Sudo = false
	expectedString = "apt-get install git-essentials vim -y -d"

	if s := command.String(); s != expectedString {
		t.Errorf(msg, spec, expectedString, s)
	}
}

func ExampleNewCommand() {
	command, err := NewCommand(true,
		"apt-get",
		"install",
		[]string{"git-essentials"},
		[]string{})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(command)
	// Output: sudo apt-get install git-essentials
}
