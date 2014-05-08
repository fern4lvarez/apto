package apto

import (
	"testing"
)

// TestUpdate is ignored when running CI since it is not
// possible to run apt-get update as su on remote hosts
func TestUpdate(t *testing.T) {
	spec := "Should return nil when Update is executed"

	if err := Update(); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}

// TestUpgrade is ignored when running CI since it is not
// possible to run apt-get dist-upgrade as su on remote hosts
func TestUpgrade(t *testing.T) {
	spec := "Should return nil when Upgrade is executed"

	if err := Upgrade(); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}
