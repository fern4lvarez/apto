package apto

import (
	"errors"
	"testing"
)

// TestInstall is ignored when running CI since it is not
// possible to run apt-get install as su on remote hosts
func IgnoreTestInstall(t *testing.T) {
	spec := "Should return nil when Install happens given correct args"
	args := []string{"apto", "install", "apt"}

	if err := Install(args); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}

func TestInstallError(t *testing.T) {
	spec := "Should return error when Install happens given wrong arguments"
	args := []string{"apto", "install"}

	if err := Install(args); err == nil {
		t.Errorf(msg, spec, err, nil)
	}
}

// TestUninstall is ignored when running CI since it is not
// possible to run apt-get install as su on remote hosts
func IgnoreTestUninstall(t *testing.T) {
	spec := "Should return nil when Install happens given correct args"
	args := []string{"apto", "uninstall", "apt"}

	if err := Uninstall(args); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}

func TestUninstallError(t *testing.T) {
	spec := "Should return error when Install happens given wrong arguments"
	args := []string{"apto", "uninstall"}

	if err := Uninstall(args); err == nil {
		t.Errorf(msg, spec, err, nil)
	}
}

func TestUnOrInstallCommand(t *testing.T) {
	spec := "Should return a new Command based on the arguments with no options"
	args := []string{"apto", "install", "vim", "git-essentials"}
	expectedCommand := "sudo apt-get install vim git-essentials -y"

	if command, err := unOrInstallCommand(args, "install"); err != nil {
		t.Errorf(msg, spec, nil, err)
	} else if commandS := command.String(); commandS != expectedCommand {
		t.Errorf(msg, spec, expectedCommand, commandS)
	}
}

func TestInstallCommandError(t *testing.T) {
	spec := "Should return an Error when no packages are given"
	args := []string{"apto", "install"}
	expectedErr := errors.New("No given pkgs to Install.")

	if _, err := unOrInstallCommand(args, "install"); err == nil {
		t.Errorf(msg, spec, expectedErr, err)
	}
}
