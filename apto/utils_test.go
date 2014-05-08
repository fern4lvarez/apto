package apto

import (
	"reflect"
	"testing"
)

func ExampleDebug() {
	a, b := 5, 18
	Debug("Value of a is %v - Value of b is %v", a, b)
	// Output: --- DEBUG: Value of a is 5 - Value of b is 18
}

func TestHandleFlag(t *testing.T) {
	spec := "Should return args without flag -f and true"
	flag := "-f"
	args := []string{"uninstall", "-f", "vim", "gnomine"}
	expectedArgs := []string{"uninstall", "vim", "gnomine"}

	if args, exists := HandleFlag(args, flag); exists == false {
		t.Errorf(msg, spec, true, exists)
	} else if !reflect.DeepEqual(expectedArgs, args) {
		t.Errorf(msg, spec, expectedArgs, args)
	}

	spec = "Should return same args when given flag not in args and return false"
	flag = "-f"
	args = []string{"uninstall", "vim", "gnomine"}
	expectedArgs = []string{"uninstall", "vim", "gnomine"}

	if args, exists := HandleFlag(args, flag); exists == true {
		t.Errorf(msg, spec, false, exists)
	} else if !reflect.DeepEqual(expectedArgs, args) {
		t.Errorf(msg, spec, expectedArgs, args)
	}
}
