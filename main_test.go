// +build integration

package main

import (
	"testing"

	"github.com/fern4lvarez/apto/apto"
)

var (
	msg = "%v. Expects %v, returns %v"
)

func TestInstall(t *testing.T) {
	spec := "Should execute `apto install gnomine`"
	command := []string{"apto", "install", "gnomine"}

	if err := apto.Execute(command); err != nil {
		t.Errorf(msg, spec, nil, err)
	}

	spec = "Should exit when install command has no packages"
	command = []string{"apto", "install"}

	if err := apto.Execute(command); err == nil {
		t.Errorf(spec)
	}
}

func TestUninstall(t *testing.T) {
	spec := "Should execute `apto uninstall gnomine`"
	command := []string{"apto", "uninstall", "gnomine"}

	if err := apto.Execute(command); err != nil {
		t.Errorf(msg, spec, nil, err)
	}

	spec = "Should execute `apto -f uninstall gnomine`"
	command = []string{"apto", "-f", "uninstall", "gnomine"}

	if err := apto.Execute(command); err != nil {
		t.Errorf(spec)
	}

	spec = "Should execute `apto uninstall -f gnomine`"
	command = []string{"apto", "uninstall", "-f", "gnomine"}

	if err := apto.Execute(command); err != nil {
		t.Errorf(spec)
	}

	spec = "Should exit when uninstall command has no packages"
	command = []string{"apto", "uninstall"}

	if err := apto.Execute(command); err == nil {
		t.Errorf(spec)
	}
}

func TestUpdate(t *testing.T) {
	spec := "Should execute `apto update`"
	command := []string{"apto", "update"}

	if err := apto.Execute(command); err != nil {
		t.Errorf(msg, spec, nil, err)
	}

	spec = "Should exit when update command has more parameters"
	command = []string{"apto", "update", "foo"}

	if err := apto.Execute(command); err == nil {
		t.Errorf(spec)
	}
}

func TestUpgrade(t *testing.T) {
	spec := "Should execute `apto upgrade`"
	command := []string{"apto", "upgrade"}

	if err := apto.Execute(command); err != nil {
		t.Errorf(msg, spec, nil, err)
	}

	spec = "Should exit when upgrade command has more parameters"
	command = []string{"apto", "upgrade", "foo"}

	if err := apto.Execute(command); err == nil {
		t.Errorf(spec)
	}
}
