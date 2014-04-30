package apto

import (
	"bufio"
	"os"
	"path/filepath"
)

var (
	home        = os.Getenv("HOME")
	current_dir = os.Getenv("PWD")
)

// Aptofile contains all parsed data from text file
type Aptofile struct {
	Commands []*Command
	Queue    chan *Command
	Lenght   int
	Location string
}

// NewAptofile creates a fresh Aptofile given a path
func NewAptofile(path string) (*Aptofile, error) {
	aptofile := new(Aptofile)
	err := aptofile.SetLocation(path)
	if err != nil {
		return nil, err
	}
	return aptofile, nil
}

// SetLocation sets location of an Aptofile given a path
func (aptofile *Aptofile) SetLocation(path string) error {
	if path == "" {
		aptofile.Location = filepath.Join(current_dir, "Aptofile")
		return nil
	}

	location := filepath.Join(path, "Aptofile")
	if _, err := os.Stat(location); err != nil {
		return err
	}

	aptofile.Location = location
	return nil
}

// Read reads and parses an Aptofile
func (aptofile *Aptofile) Read() error {
	file, err := os.Open(aptofile.Location)
	if err != nil {
		return err
	}

	commands := []*Command{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := NewCommand()
		command.handleLine(scanner.Text())
		commands = command.appendTo(commands)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	aptofile.Commands = commands

	return nil
}

// Execute executes the aptofile command by command
func (aptofile *Aptofile) Execute() error {
	for _, command := range aptofile.Commands {
		command.Execute()
	}

	return nil
}

// File reads and executes an aptofile
func File(args []string) {
	aptofile, _ := NewAptofile("")
	aptofile.Read()
	aptofile.Execute()
}
