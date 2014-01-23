package apto

import (
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
	} else {
		_, err := os.Stat(path)
		if err != nil {
			return err
		}
		aptofile.Location = path
	}
	return nil
}

func Bundle(args []string) {
	fmt.Println("Bundle!!")
}
