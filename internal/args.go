package internal

import (
	"fmt"
	"os"
	"slices"
)

func Args() (host string) {
	args := os.Args[1:]
	help := slices.Contains(args, "help")
	if len(args) < 1 {
		fmt.Println("Error, you need to indicate a host to connect to")
		fmt.Println("Usage:")
		fmt.Println("       ./intr user@host")
		return
	}
	if len(args) > 1 {
		fmt.Println("Error, you can only indicate one host to connect to")
		return
	}
	if help {
		fmt.Println("Usage:")
		fmt.Println("       ./intr user@host")
		fmt.Println("")
		return
	}
	if len(args) == 1 {
		host = args[0]
		fmt.Println("You are connected to: " + host)
	}
	return host
}
