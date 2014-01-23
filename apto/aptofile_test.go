package apto

import (
	"reflect"
	"testing"
)

func TestNewAptofile(t *testing.T) {
	spec := "Should create an Aptofile with default location"
	expectedAptofile := new(Aptofile)
	expectedAptofile.Location = home

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
	spec := "Should set $HOME as Aptofile directory when given path is empty"
	expectedLocation := home
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

func TestAptofileSetLocationError(t *testing.T) {
	spec := "Should return error when given path does not exist"
	aptofile := &Aptofile{}
	path := "foobar"

	if err := aptofile.SetLocation(path); err == nil {
		t.Errorf(msg, spec, err, nil)
	}
}
