package apto

import (
	"errors"
	"testing"
)

func TestInstallError(t *testing.T) {
	spec := "Should return error when Install happens given wrong arguments"
	args := []string{"install"}

	if err := Install(args); err == nil {
		t.Errorf(msg, spec, err, nil)
	}
}

func TestUninstallError(t *testing.T) {
	spec := "Should return error when Install happens given wrong arguments"
	args := []string{"uninstall"}

	if err := Uninstall(args, false); err == nil {
		t.Errorf(msg, spec, err, nil)
	}
}

func TestUnOrInstallCommand(t *testing.T) {
	spec := "Should return a new Command based on the arguments with no options"
	args := []string{"install", "vim", "git-essentials"}
	expectedCommand := "sudo apt-get install vim git-essentials -y"

	if command, err := unOrInstallCommand(args, "install", false); err != nil {
		t.Errorf(msg, spec, nil, err)
	} else if commandS := command.String(); commandS != expectedCommand {
		t.Errorf(msg, spec, expectedCommand, commandS)
	}
}

func TestInstallCommandError(t *testing.T) {
	spec := "Should return an Error when no packages are given"
	args := []string{"install"}
	expectedErr := errors.New("No given pkgs to Install.")

	if _, err := unOrInstallCommand(args, "install", false); err == nil {
		t.Errorf(msg, spec, expectedErr, err)
	}
}
