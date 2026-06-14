package main

import (
	"strings"
	"testing"
)

func TestInputCommand(t *testing.T) {
	// Create a "fake" input buffer
	input := "ls -la\n"
	r := strings.NewReader(input)

	got := InputCommand(r)
	want := "ls -la"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestValidateCommand(t *testing.T) {
	for _, c := range List {
		if !ValidateCommand(c, List) {
			t.Errorf("Command not found %v", c)
		}
	}

	invalidCmd := "invalid_command"
	if ValidateCommand(invalidCmd, List) {
		t.Errorf("Expected %s to be invalid, but it passed!", invalidCmd)
	}
}

func TestSplitCommand(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Standard command", "ls -lisah", "ls"},
		{"No arguments", "top", "top"},
		{"Leading spaces", "   grep pattern", "grep"},
		{"Multiple spaces", "sh  -c  'echo hi'", "sh"},
		{"Empty input", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitCommand(tt.input)
			if got != tt.want {
				t.Errorf("SplitCommand(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
