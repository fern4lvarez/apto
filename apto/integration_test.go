// +build integration

package apto

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	spec := "Should return nil when executing a command"
	command := []string{"echo", "hello"}

	if err := Execute(command); err != nil {
		t.Errorf(msg, spec, nil, err)
	}

}

func TestInstall(t *testing.T) {
	spec := "Should return nil when Install happens given correct args"
	args := []string{"install", "gnomine"}

	if err := Install(args); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}

func TestUninstall(t *testing.T) {
	spec := "Should return nil when Uninstall is executed given correct args"
	args := []string{"uninstall", "gnomine"}
	force := false

	if err := Uninstall(args, force); err != nil {
		t.Errorf(msg, spec, nil, err)
	}

	spec = "Should return nil when force Uninstall happens given correct args"
	args = []string{"uninstall", "gnomine"}
	force = true

	if err := Uninstall(args, force); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}

func TestUpdate(t *testing.T) {
	spec := "Should return nil when Update is executed"

	if err := Update(); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}

func TestUpgrade(t *testing.T) {
	spec := "Should return nil when Upgrade is executed"

	if err := Upgrade(); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}

func TestAptofileExecute(t *testing.T) {
	spec := `Should execute Aptofile`

	af := []byte("upgrade\ninstall vim\ninstall gnomine\nremove -f gnomine\n")
	ioutil.WriteFile("Aptofile", af, 0644)
	defer os.Remove("Aptofile")
	aptofile, _ := NewAptofile(current_dir)
	aptofile.Read()

	if err := aptofile.Execute(); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}

func TestAptofileFile(t *testing.T) {
	spec := "Should read and execute Aptofile from current path"

	af := []byte("upgrade\ninstall vim\ninstall gnomine\nremove -f gnomine\n")
	ioutil.WriteFile("Aptofile", af, 0644)
	defer os.Remove("Aptofile")

	if err := File([]string{}); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}
