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
	expectedAptofile.Location = filepath.Join(current_dir, "Aptofile")

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
	spec := "Should set $PWD/Aptofile as Aptofile directory when given path is empty"
	ioutil.WriteFile("Aptofile", []byte(""), 0644)
	defer os.Remove("Aptofile")
	expectedLocation := filepath.Join(current_dir, "Aptofile")
	aptofile := &Aptofile{}

	if err := aptofile.SetLocation(""); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
	if location := aptofile.Location; location != expectedLocation {
		t.Errorf(msg, spec, expectedLocation, location)
	}

	spec = "Should set as Aptofile directory the given path"
	aptofile.Location = ""
	path := current_dir
	expectedLocation = filepath.Join(path, "Aptofile")

	if err := aptofile.SetLocation(path); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
	if location := aptofile.Location; location != expectedLocation {
		t.Errorf(msg, spec, expectedLocation, location)
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

func TestAptofileRead(t *testing.T) {
	spec := `Should complete all Aptofile after reading
file with two install commands`

	af := []byte("install vim\ninstall gnomine\n")
	ioutil.WriteFile("Aptofile", af, 0644)

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
	expectedAptofile, _ := NewAptofile(current_dir)
	expectedAptofile.Commands = []*Command{command1, command2}
	aptofile, _ := NewAptofile(current_dir)

	if err := aptofile.Read(); err != nil {
		t.Errorf(msg, spec, nil, err)
	}

	if !reflect.DeepEqual(expectedAptofile, aptofile) {
		t.Errorf(msg, spec, expectedAptofile, aptofile)
	}

	os.Remove("Aptofile")

	spec = `Should return an Aptofile with an empty list of commands
    given empty lines or wrong commands`

	af = []byte("\n\n\n\nwrong command\n\n\nevil command\n")
	ioutil.WriteFile("Aptofile", af, 0644)
	defer os.Remove("Aptofile")

	expectedAptofile, _ = NewAptofile(current_dir)
	expectedAptofile.Commands = []*Command{}
	aptofile, _ = NewAptofile(current_dir)

	if err := aptofile.Read(); err != nil {
		t.Errorf(msg, spec, nil, err)
	}

	if !reflect.DeepEqual(expectedAptofile, aptofile) {
		t.Errorf(msg, spec, expectedAptofile, aptofile)
	}
}

func TestAptofileReadError(t *testing.T) {
	spec := `Should return error if Aptofile's location is not existing`

	ioutil.WriteFile("Aptofile", []byte("foo"), 0644)
	expectedErr := "no such file or directory"
	aptofile, _ := NewAptofile(current_dir)

	os.Remove("Aptofile")
	if err := aptofile.Read(); err == nil {
		t.Errorf(msg, spec, expectedErr, err)
	}
}
