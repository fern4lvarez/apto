package apto

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	home = os.Getenv("HOME")
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
		aptofile.Location = home
		return nil
	}

	if _, err := os.Stat(path); err != nil {
		return err
	}
	aptofile.Location = path
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
		line := strings.TrimSpace(scanner.Text())
		args := strings.Split(line, " ")
		switch cmd := args[0]; cmd {
		case "install":
			command.Install(args[1:], []string{})
		default:
			continue
		}

		commands = append(commands, command)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	aptofile.Commands = commands

	return nil
}

func Bundle(args []string) {
	fmt.Println("Bundle!!")
}
