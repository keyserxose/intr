package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var list = []string{
	"cat",
	"dd",
	"echo",
	"test",
	"tree",
	"tail",
	"cp",
	"ls",
	"mkdir",
	"pwd",
	"chmod",
	"ln",
	"mv",
	"rm",
	"rmdir",
	"touch",
	"chgrp",
	"groups",
	"id",
	"passwd",
	"quota",
	"find",
	"gfind",
	"split",
	"fetch",
	"df",
	"md5",
	"sha1",
	"sha224",
	"sha256",
	"sha384",
	"sha512",
	"sha512t256",
	"rmd160",
	"skein256",
	"skein512",
	"skein1024",
	"date",
	"cksum",
}

var exitCommands = []string{
	"exit",
	"quit",
	"break",
}

func main() {
	host := args()
	for {
		command := inputCommand()
		validateCommand(host, command)
	}
}

func inputCommand() (command string) {
	fmt.Print("Input Command: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	command = input.Text()
	fmt.Println("")
	return command
}

func validateCommand(host string, command string) {
	singleCommand := splitCommand(command)
	availableCommand := slices.Contains(list, singleCommand)
	if availableCommand {
		terminal(host, command)
	}
	if !availableCommand {
		availableExitCommand := slices.Contains(exitCommands, command)
		if availableExitCommand {
			os.Exit(1)
		} else {
			fmt.Println("Wrong command!")
			fmt.Println("")
			fmt.Println("These are the available commands: " + strings.Join(list, " "))
		}
	}
}

// This is used is to remove the parameters of a command in order to validate it
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

func args() (host string) {
	args := os.Args[1:]
	help := slices.Contains(args, "help")
	if len(args) < 1 {
		fmt.Println("Error, you need to indicate a host to connect to")
		fmt.Println("Usage:")
		fmt.Println("       ./rsync user@host")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("Error, you can only indicate one host to connect to")
		os.Exit(1)
	}
	if help {
		fmt.Println("Usage:")
		fmt.Println("       ./rsync user@host")
		fmt.Println("")
		os.Exit(1)
	}
	if len(args) == 1 {
		host = args[0]
		fmt.Println("You are connected to: " + host)
	}
	return host
}
