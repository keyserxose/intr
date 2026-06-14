package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	host := Args()
	err := PreFlightCheck(host)
	if err != nil {
		// This is where you will finally see WHY it was failing
		fmt.Printf("Execution failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("You are connected to: " + host)
	for {
		command := InputCommand(os.Stdin)
		if ValidateCommand(command, ExitCommands) {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		if !ValidateCommand(command, List) {
			fmt.Println("Please, enter a valid command")
			fmt.Println("")
			fmt.Println("These are the available commands: " + strings.Join(List, " "))
			fmt.Println("")
			continue
		}
		err := RunCommand(host, command)
		if err != nil {
			// This is where you will finally see WHY it was failing
			fmt.Printf("Execution failed: %v\n", err)
		}
	}
}
