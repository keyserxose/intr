// utility to run commands against rsync.net
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

var host string
var command string
var singleCommand string
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
	args()
	for {
		inputCommand()
		validateCommand()
	}
}

func inputCommand() {
	fmt.Print("Input Command: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	command = input.Text()
	fmt.Println("")
}

func validateCommand() {
	splitCommand()
	availableCommand := slices.Contains(list, singleCommand)
	if availableCommand == true {
		terminal()
	}
	if availableCommand == false {
		availableExitCommand := slices.Contains(exitCommands, command)
		if availableExitCommand == true {
			os.Exit(1)
		} else {
			fmt.Println("Wrong command!")
			fmt.Println("")
			fmt.Println("These are the available commands: " + strings.Join(list, " "))
		}
	}
}

func splitCommand() {
	commandsSeparated := strings.Split(command, " ")
	singleCommand = commandsSeparated[0]
}

func terminal() {
	cmd := exec.Command("ssh", host, command)
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.Stdout)
}

/* func flags_disabled() {
	host = flag.String("host", "", "Indicate the hostname")
	flag.Parse()
	hostStr := *host
	if hostStr == "" {
		fmt.Println("Error, you need to indicate a host to connect to")
		os.Exit(1)
	}
} */

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func args() {
	args := os.Args[1:]
	help := contains(args, "help")
	if len(args) < 1 {
		fmt.Println("Error, you need to indicate a host to connect to")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("Error, you can only indicate one host")
		os.Exit(1)
	}
	if help == true {
		fmt.Println("Usage:")
		fmt.Println("       ./rsync user@host")
		fmt.Println("")
		os.Exit(1)
	}
	if len(args) == 1 {
		host = args[0]
		fmt.Println("You are connected to: " + host)
	}
}
