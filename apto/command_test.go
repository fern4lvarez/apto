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

func IgnoreTestExecute(t *testing.T) {
	spec := "Should execute not sudo bash command and return output"
	command := &Command{Sudo: false,
		Tool:    "apt-get",
		Cmd:     "help",
		Pkgs:    []string{},
		Options: []string{}}

	if err := command.Execute(); err != nil {
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

	if err := command.Execute(); err == nil {
		t.Errorf(msg, spec, err, nil)
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
	expectedErr := errors.New("No given pkgs to install.")

	if err := command.Install(
		[]string{},
		[]string{}); err == nil {
		t.Errorf(msg, spec)
	} else if !reflect.DeepEqual(expectedErr, err) {
		t.Errorf(msg, spec, expectedErr, err)
	}
}

func TestCommandUninstall(t *testing.T) {
	spec := "Should create an Uninstall command given pkgs and options"
	command := NewCommand()
	expectedCommand := Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "remove",
		Pkgs:    []string{"git-essentials"},
		Options: []string{"-y"},
	}

	if err := command.Uninstall(
		[]string{"git-essentials"},
		[]string{},
		false); err != nil {
		t.Errorf(msg, spec, err, nil)
	}

	if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}

	spec = "Should create an forced Uninstall command given pkgs and options"
	command = NewCommand()
	expectedCommand = Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "purge",
		Pkgs:    []string{"git-essentials"},
		Options: []string{"-y"},
	}

	if err := command.Uninstall(
		[]string{"git-essentials"},
		[]string{},
		true); err != nil {
		t.Errorf(msg, spec, err, nil)
	}

	if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}
}

func TestCommandUninstallError(t *testing.T) {
	spec := "Should return error when not given pkgs"
	command := NewCommand()
	expectedErr := errors.New("No given pkgs to remove.")

	if err := command.Uninstall(
		[]string{},
		[]string{},
		false); err == nil {
		t.Errorf(msg, spec)
	} else if !reflect.DeepEqual(expectedErr, err) {
		t.Errorf(msg, spec, expectedErr, err)
	}
}

func TestCommandUpdate(t *testing.T) {
	spec := "Should create an Update command"
	command := NewCommand()
	expectedCommand := Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "update",
		Pkgs:    []string{},
		Options: []string{},
	}

	command.Update()

	if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}
}

func TestCommandUpgrade(t *testing.T) {
	spec := "Should create an upgrade command"
	command := NewCommand()
	expectedCommand := Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "dist-upgrade",
		Pkgs:    []string{},
		Options: []string{"-y"},
	}

	command.Upgrade()

	if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}
}

func TestCommandShell(t *testing.T) {
	spec := "Should create shell command given instructions"
	instructions := []string{"bundle", "install", "--verbose"}
	command := NewCommand()
	expectedCommand := Command{Sudo: false,
		Tool:    "bundle",
		Cmd:     "install",
		Pkgs:    []string{"--verbose"},
		Options: []string{},
	}

	if err := command.Shell(instructions); err != nil {
		t.Errorf(msg, spec, err, nil)
	}

	if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}
}

func TestCommandShellError(t *testing.T) {
	spec := "Should return error when instructions are empty"
	command := NewCommand()
	expectedErr := errors.New("No given instructions.")

	if err := command.Shell([]string{}); err == nil {
		t.Errorf(msg, spec)
	} else if !reflect.DeepEqual(expectedErr, err) {
		t.Errorf(msg, spec, expectedErr, err)
	}
}

func TestCommandEcho(t *testing.T) {
	spec := "Should create echo command given a text"
	text := "This is a comment"
	command := NewCommand()
	expectedCommand := Command{Sudo: false,
		Tool:    "echo",
		Cmd:     text,
		Pkgs:    []string{},
		Options: []string{},
	}

	if err := command.Echo(text); err != nil {
		t.Errorf(msg, spec, err, nil)
	}

	if !reflect.DeepEqual(expectedCommand, *command) {
		t.Errorf(msg, spec, expectedCommand, *command)
	}
}

func TestCommandEchoError(t *testing.T) {
	spec := "Should return error when text is empty"
	command := NewCommand()
	expectedErr := errors.New("Empty text.")

	if err := command.Echo(""); err == nil {
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
	spec := "Should convert a Echo command given comment"
	expectedCommand := &Command{Sudo: false,
		Tool:    "echo",
		Cmd:     "this is a comment",
		Pkgs:    []string{},
		Options: []string{},
	}
	line := "# this is a comment"
	command := NewCommand()
	command.handleLine(line)

	if !reflect.DeepEqual(expectedCommand, command) {
		t.Errorf(msg, spec, expectedCommand, command)
	}

	spec = "Should convert an Install command given a install line"
	expectedCommand = &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"git-essentials"},
		Options: []string{"-y"},
	}
	line = "install git-essentials"
	command = NewCommand()
	command.handleLine(line)

	if !reflect.DeepEqual(expectedCommand, command) {
		t.Errorf(msg, spec, expectedCommand, command)
	}

	spec = "Should convert an Uninstall command given a uninstall line"
	expectedCommand = &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "remove",
		Pkgs:    []string{"git-essentials"},
		Options: []string{"-y"},
	}
	line = "uninstall git-essentials"
	command = NewCommand()
	command.handleLine(line)

	if !reflect.DeepEqual(expectedCommand, command) {
		t.Errorf(msg, spec, expectedCommand, command)
	}

	spec = "Should convert a forced Uninstall command given a uninstall line with force flag"
	expectedCommand = &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "purge",
		Pkgs:    []string{"git-essentials"},
		Options: []string{"-y"},
	}
	line = "uninstall -f git-essentials"
	command = NewCommand()
	command.handleLine(line)

	if !reflect.DeepEqual(expectedCommand, command) {
		t.Errorf(msg, spec, expectedCommand, command)
	}

	spec = "Should convert an Update command given an update line"
	expectedCommand = &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "update",
		Pkgs:    []string{},
		Options: []string{},
	}
	line = "update"
	command = NewCommand()
	command.handleLine(line)

	spec = "Should convert an Upgrade command given an upgrade line"
	expectedCommand = &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "dist-upgrade",
		Pkgs:    []string{},
		Options: []string{"-y"},
	}
	line = "upgrade"
	command = NewCommand()
	command.handleLine(line)
	if !reflect.DeepEqual(expectedCommand, command) {
		t.Errorf(msg, spec, expectedCommand, command)
	}

	spec = "Should convert a Shell command given a shell line"
	expectedCommand = &Command{Sudo: false,
		Tool:    "exec",
		Cmd:     "whatever",
		Pkgs:    []string{"command", "you", "want"},
		Options: []string{},
	}
	line = "$ exec whatever command you want"
	command = NewCommand()
	command.handleLine(line)

	if !reflect.DeepEqual(expectedCommand, command) {
		t.Errorf(msg, spec, expectedCommand, command)
	}
}

func TestCommandHandleLineError(t *testing.T) {
	spec := "Should set an Empty command given an empty line"
	expectedCommand := NewCommand()
	line := ""
	command := NewCommand()
	command.handleLine(line)

	if !reflect.DeepEqual(expectedCommand, command) {
		t.Errorf(msg, spec, expectedCommand, command)
	}

	spec = "Should set an empty Command given a supported command"
	expectedCommand = NewCommand()
	line = "wrong command"
	command = NewCommand()
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
