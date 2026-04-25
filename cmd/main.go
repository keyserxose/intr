package main

import "github.com/keyserxose/intr/internal"

func main() {
	host := internal.Args()
	for {
		command := internal.InputCommand()
		internal.ValidateCommand(host, command)
	}
}
