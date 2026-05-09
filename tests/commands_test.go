package tests

import (
	"testing"

	"github.com/keyserxose/intr/internal"
)

func TestValidateCommand(t *testing.T) {
	for _, c := range internal.List {
		if !internal.ValidateCommand(c, internal.List) {
			t.Errorf("Command not found %v", c)
		}
	}

	invalidCmd := "invalid_command"
	if internal.ValidateCommand(invalidCmd, internal.List) {
		t.Errorf("Expected %s to be invalid, but it passed!", invalidCmd)
	}
}
