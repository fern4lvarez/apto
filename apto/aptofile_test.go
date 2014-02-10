package apto

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewAptofile(t *testing.T) {
	spec := "Should create an Aptofile with default location"
	expectedAptofile := new(Aptofile)
	expectedAptofile.Location = filepath.Join(home, "Aptofile")

	if aptofile, err := NewAptofile(""); err != nil {
		t.Errorf(msg, spec, nil, err)
	} else if !reflect.DeepEqual(expectedAptofile, aptofile) {
		t.Errorf(msg, spec, expectedAptofile, aptofile)
	}
}

func TestNewAptofileError(t *testing.T) {
	spec := "Should return an error when creation a new Aptofile with a wrong path"
	path := "foobar"

	if _, err := NewAptofile(path); err == nil {
		t.Errorf(msg, spec, err, nil)
	}
}

func TestAptofileSetLocation(t *testing.T) {
	spec := "Should set $HOME/Aptofile as Aptofile directory when given path is empty"
	expectedLocation := filepath.Join(home, "Aptofile")
	aptofile := &Aptofile{}

	if err := aptofile.SetLocation(""); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
	if location := aptofile.Location; location != expectedLocation {
		t.Errorf(msg, spec, expectedLocation, location)
	}

	spec = "Should set as Aptofile directory the given path"
	aptofile.Location = ""
	expectedLocation = home
	path := expectedLocation

	if err := aptofile.SetLocation(path); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
	if location := aptofile.Location; location != expectedLocation {
		t.Errorf(msg, spec, expectedLocation, location)
	}
}

func TestAptofileRead(t *testing.T) {
	spec := `Should complete all Aptofile after reading
file with two install commands`

	af := []byte("install vim\ninstall gnomine\n")
	ioutil.WriteFile("Aptofile", af, 0644)
	defer os.Remove("Aptofile")

	command1 := &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"vim"},
		Options: []string{"-y"},
	}
	command2 := &Command{Sudo: true,
		Tool:    "apt-get",
		Cmd:     "install",
		Pkgs:    []string{"gnomine"},
		Options: []string{"-y"},
	}
	expectedAptofile, _ := NewAptofile("Aptofile")
	expectedAptofile.Commands = []*Command{command1, command2}
	aptofile, _ := NewAptofile("Aptofile")

	if err := aptofile.Read(); err != nil {
		t.Errorf(msg, spec, nil, err)
	} else if !reflect.DeepEqual(expectedAptofile, aptofile) {
		t.Errorf(msg, spec, expectedAptofile, aptofile)
	}

	spec = `Should return an Aptofile with an empty list of commands
    given empty lines or wrong commands`

	af = []byte("\n\n\n\nwrong command\n\n\nevil command\n")
	ioutil.WriteFile("Aptofile2", af, 0644)
	defer os.Remove("Aptofile2")

	expectedAptofile, _ = NewAptofile("Aptofile2")
	expectedAptofile.Commands = []*Command{}
	aptofile, _ = NewAptofile("Aptofile2")

	if err := aptofile.Read(); err != nil {
		t.Errorf(msg, spec, nil, err)
	} else if !reflect.DeepEqual(expectedAptofile, aptofile) {
		t.Errorf(msg, spec, expectedAptofile, aptofile)
	}
}

func TestAptofileSetLocationError(t *testing.T) {
	spec := "Should return error when given path does not exist"
	aptofile := &Aptofile{}
	path := "foobar"

	if err := aptofile.SetLocation(path); err == nil {
		t.Errorf(msg, spec, err, nil)
	}
}
