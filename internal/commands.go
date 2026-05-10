package internal

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"slices"
	"strings"
)

func InputCommand(r io.Reader) string {
	input := bufio.NewScanner(r)
	input.Scan()
	command := input.Text()
	fmt.Println("")
	return command
}

func ValidateCommand(command string, list []string) bool {
	singleCommand := SplitCommand(command)
	return slices.Contains(list, singleCommand)
}

// This is used to remove the parameters of a command in order to validate it
func SplitCommand(command string) string {
	commandsSeparated := strings.Fields(command)
	if len(commandsSeparated) == 0 {
		return ""
	}
	return commandsSeparated[0]
}

func RunCommand(host string, command string) error {
	cmd := exec.Command("ssh", host, command)

	var stdout strings.Builder
	var stderr strings.Builder

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("command failed: %w, stderr: %s", err, stderr.String())
	}
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())
	return nil
}

func PreFlightCheck(host string) error {
	cmd := exec.Command("ssh", host)

	var stderr strings.Builder

	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("command failed: %w, stderr: %s", err, stderr.String())
	}
	return nil

}
