package main

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
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("Error, you can only indicate one host to connect to")
		os.Exit(1)
	}
	if help {
		fmt.Println("Usage:")
		fmt.Println("       ./intr user@host")
		fmt.Println("")
		os.Exit(0)
	}
	if len(args) == 1 {
		host = args[0]
	}
	return host
}
