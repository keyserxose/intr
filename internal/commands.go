package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func InputCommand() (command string) {
	//fmt.Print("Input Command: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	command = input.Text()
	fmt.Println("")
	return command
}

func ValidateCommand(host string, command string) {
	singleCommand := splitCommand(command)
	availableCommand := slices.Contains(list, singleCommand)
	if availableCommand {
		terminal(host, command)
	}
	if !availableCommand {
		availableExitCommand := slices.Contains(exitCommands, command)
		if availableExitCommand {
			os.Exit(0)
		} else {
			fmt.Println("Wrong command!")
			fmt.Println("")
			fmt.Println("These are the available commands: " + strings.Join(list, " "))
		}
	}
}

// This is used to remove the parameters of a command in order to validate it
func splitCommand(command string) (singleCommand string) {
	commandsSeparated := strings.Split(command, " ")
	singleCommand = commandsSeparated[0]
	return singleCommand
}

func terminal(host string, command string) {
	cmd := exec.Command("ssh", host, command)
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.Stdout)
}
