package apto

import (
	"testing"
)

// TestUpdate is ignored when running CI since it is not
// possible to run apt-get update as su on remote hosts
func IgnoreTestUpdate(t *testing.T) {
	spec := "Should return nil when Update is executed"

	if err := Update(); err != nil {
		t.Errorf(msg, spec, nil, err)
	}
}
