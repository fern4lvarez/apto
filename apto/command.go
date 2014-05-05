/*
apto is a package automatically generated by ´gobi´. Happy hacking!
*/
package apto

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// MyaptoExample is a example type automatically generated by ´gobi´.
type Command struct {
	Sudo    bool
	Tool    string
	Cmd     string
	Pkgs    []string
	Options []string
}

// NewCommand creates new, empty *Command
func NewCommand() *Command {
	return new(Command)
}

// Create *Command given arguments
func (command *Command) Create(sudo bool, tool string, cmd string, pkgs []string, options []string) error {
	if tool == "" {
		return errors.New("Tool is empty.")
	}

	if cmd == "" {
		return errors.New("Cmd is empty.")
	}

	command.Sudo = sudo
	command.Tool = tool
	command.Cmd = cmd
	command.Pkgs = pkgs
	command.Options = options

	return nil
}

// Execute executes Command
func (command *Command) Execute() error {
	c := strings.Split(command.String(), " ")
	cmd := exec.Command(c[0], c[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Install creates a apt-get install command given packages and options
func (command *Command) Install(pkgs []string, options []string) error {
	return command.UnOrInstall(pkgs, options, "install")
}

// Uninstall creates a apt-get remove command given packages and options
func (command *Command) Uninstall(pkgs []string, options []string, force bool) error {
	method := "remove"

	if force {
		method = "purge"
	}

	return command.UnOrInstall(pkgs, options, method)
}

// Shell creates a shell command given instructions
func (command *Command) Shell(instructions []string) error {
	if len(instructions) == 0 {
		return errors.New("No given instructions.")
	}

	command.Create(false,
		instructions[0],
		instructions[1],
		instructions[2:],
		[]string{})

	return nil
}

// Echo creates a echo command given a text
func (command *Command) Echo(text string) error {
	if text == "" {
		return errors.New("Empty text.")
	}

	command.Create(false,
		"echo",
		text,
		[]string{},
		[]string{})

	return nil
}

// String ressembles Command into a valid bash command
func (command *Command) String() string {
	var buf bytes.Buffer
	if command.Sudo == true {
		buf.WriteString("sudo ")
	}

	buf.WriteString(command.Tool)
	buf.WriteString(" ")
	buf.WriteString(command.Cmd)
	buf.WriteString(" ")

	for i := range command.Pkgs {
		buf.WriteString(command.Pkgs[i])
		buf.WriteString(" ")
	}

	for i := range command.Options {
		buf.WriteString(command.Options[i])
		buf.WriteString(" ")
	}

	return strings.TrimSpace(buf.String())
}

// handleLine turns a Command given a string line
func (command *Command) handleLine(line string) {
	line = strings.TrimSpace(line)

	if line == "" {
		return
	}

	if strings.HasPrefix(line, "#") {
		command.Echo(strings.Trim(line, "# "))
		return
	}

	args := strings.Split(line, " ")

	switch cmd := args[0]; cmd {
	case "install":
		command.Install(args[1:], []string{})
	case "uninstall":
		command.Uninstall(args[1:], []string{}, false)
	case "$":
		command.Shell(args[1:])
	default:
		command = NewCommand()
	}
}

// unOrInstall create an install or uninstall command given the method name
func (command *Command) UnOrInstall(pkgs []string, options []string, method string) error {
	if len(pkgs) == 0 {
		return errors.New(fmt.Sprintf("No given pkgs to %s.", method))
	}

	options = append(options, "-y")

	command.Create(true,
		"apt-get",
		method,
		pkgs,
		options)

	return nil
}

// appendTo a given list of commands if command is empty
func (command *Command) appendTo(commands []*Command) []*Command {
	if command.isValid() {
		return append(commands, command)
	}

	return commands
}

// isValid checks if command is valid for using purposes
func (command *Command) isValid() bool {
	if command == nil {
		return false
	}

	if command.Tool == "" || command.Cmd == "" {
		return false
	}

	return true
}
