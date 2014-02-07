package apto

import (
	"bufio"
	"fmt"
	"os"
)

var (
	home = os.Getenv("HOME")
)

type Aptofile struct {
	Commands []*Command
	Queue    chan *Command
	Lenght   int
	Location string
}

func NewAptofile(path string) (*Aptofile, error) {
	aptofile := new(Aptofile)
	err := aptofile.SetLocation(path)
	if err != nil {
		return nil, err
	}
	return aptofile, nil
}

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

func (aptofile *Aptofile) Read() error {
	file, err := os.Open(aptofile.Location)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func Bundle(args []string) {
	fmt.Println("Bundle!!")
}
