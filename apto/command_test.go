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
	} else if !reflect.DeepEqual(expectedCommand, *command) {
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
		[]string{"-y"}); err != nil {
		t.Errorf(msg, spec, err, nil)
	} else if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}
}

func TestCommandInstallError(t *testing.T) {
	spec := "Should return error when not given pkgs"
	command := NewCommand()
	expectedErr := errors.New("No given pkgs to Install.")

	if err := command.Install(
		[]string{},
		[]string{"-y"}); err == nil {
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
