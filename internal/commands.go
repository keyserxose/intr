package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func InputCommand() (command string) {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	command = input.Text()
	fmt.Println("")
	return command
}

func ValidateCommand(command string, list []string) (validation bool) {
	singleCommand := splitCommand(command)
	availableCommand := slices.Contains(list, singleCommand)
	if availableCommand {
		return true
	}
	if !availableCommand {
		return false
	}
	return validation
}

// This is used to remove the parameters of a command in order to validate it
func splitCommand(command string) (singleCommand string) {
	commandsSeparated := strings.Split(command, " ")
	singleCommand = commandsSeparated[0]
	return singleCommand
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
