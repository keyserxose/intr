package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/keyserxose/intr/internal"
)

func main() {
	host := internal.Args()
	err := internal.PreFlightCheck(host)
	if err != nil {
		// This is where you will finally see WHY it was failing
		fmt.Printf("Execution failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("You are connected to: " + host)
	for {
		command := internal.InputCommand()
		if internal.ValidateCommand(command, internal.ExitCommands) {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		if !internal.ValidateCommand(command, internal.List) {
			fmt.Println("Please, enter a valid command")
			fmt.Println("")
			fmt.Println("These are the available commands: " + strings.Join(internal.List, " "))
			fmt.Println("")
			continue
		}
		err := internal.RunCommand(host, command)
		if err != nil {
			// This is where you will finally see WHY it was failing
			fmt.Printf("Execution failed: %v\n", err)
		}
	}
}
