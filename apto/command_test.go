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

func TestCommandCreate(t *testing.T) {
	spec := "Should create a Command given all arguments"
	command := NewCommand()
	expectedCommand := Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"git-essentials"},
		Options: []string{},
	}

	if err := command.Create(true,
		"apt-get",
		"install",
		[]string{"git-essentials"},
		[]string{}); err != nil {
		t.Errorf(msg, spec, nil, err)
	}

	if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}
}

func TestCommandCreateError(t *testing.T) {
	spec := "Should return error when Tool is empty"
	command := NewCommand()
	expectedErr := errors.New("Tool is empty.")
	if err := command.Create(true,
		"",
		"install",
		[]string{"git-essentials"},
		[]string{}); !reflect.DeepEqual(expectedErr, err) {
		t.Errorf(msg, spec, expectedErr, err)
	}

	spec = "Should return error when Cmd is empty"
	expectedErr = errors.New("Cmd is empty.")
	if err := command.Create(true,
		"apt-get",
		"",
		[]string{"git-essentials"},
		[]string{}); !reflect.DeepEqual(expectedErr, err) {
		t.Errorf(msg, spec, expectedErr, err)
	}
}

func TestCommandInstall(t *testing.T) {
	spec := "Should create Install command given pkgs and options"
	command := NewCommand()
	expectedCommand := Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"git-essentials"},
		Options: []string{"-y"},
	}

	if err := command.Install(
		[]string{"git-essentials"},
		[]string{}); err != nil {
		t.Errorf(msg, spec, err, nil)
	}

	if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}
}

func TestCommandInstallError(t *testing.T) {
	spec := "Should return error when not given pkgs"
	command := NewCommand()
	expectedErr := errors.New("No given pkgs to Install.")

	if err := command.Install(
		[]string{},
		[]string{}); err == nil {
		t.Errorf(msg, spec)
	} else if !reflect.DeepEqual(expectedErr, err) {
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

func TestCommandHandleLine(t *testing.T) {
	spec := "Should change the command according to the given line"
	expectedCommand := &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"git-essentials"},
		Options: []string{"-y"},
	}
	line := "install git-essentials"
	command := NewCommand()
	command.handleLine(line)

	if !reflect.DeepEqual(expectedCommand, command) {
		t.Errorf(msg, spec, expectedCommand, command)
	}
}

func TestCommandHandleLineError(t *testing.T) {
	spec := "Should set command as when not supported command"
	expectedCommand := NewCommand()
	line := "wrong command"
	command := NewCommand()
	command.handleLine(line)

	if !reflect.DeepEqual(expectedCommand, command) {
		t.Errorf(msg, spec, expectedCommand, command)
	}
}

func TestCommandAppendTo(t *testing.T) {
	spec := "Should append to given list of commands if valid command"
	command := &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"git-essentials", "vim"},
		Options: []string{"-y", "-d"},
	}
	expectedCommands := command.appendTo([]*Command{})

	if !reflect.DeepEqual(expectedCommands[0], command) {
		t.Errorf(msg, spec, expectedCommands[0], command)
	}

	spec = "Should not append to given list if not valid command"
	command = NewCommand()
	expectedCommands = command.appendTo(expectedCommands)

	if l := len(expectedCommands); l == 2 {
		t.Errorf(msg, spec, 1, l)
	}
}

func TestCommandIsValid(t *testing.T) {
	spec := "Should return true given a valid command"
	command := &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"git-essentials", "vim"},
		Options: []string{"-y", "-d"},
	}

	if v := command.isValid(); !v {
		t.Errorf(msg, spec, true, v)
	}

	spec = "Should return false given a new command"
	command = NewCommand()

	if v := command.isValid(); v {
		t.Errorf(msg, spec, false, v)
	}

	spec = "Should return false given nil"
	command = nil

	if v := command.isValid(); v {
		t.Errorf(msg, spec, false, v)
	}
}

func ExampleCommandNewCommand() {
	command := NewCommand()
	err := command.Create(true,
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
