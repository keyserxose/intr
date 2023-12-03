// utility to run commands against rsync.net

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var host *string
var input string
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
	cmd := exec.Command("ssh", *host, command)
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.Stdout)
}

func args() {
	host = flag.String("host", "", "Indicate the hostname")
	flag.Parse()
	hostStr := *host
	if hostStr == "" {
		fmt.Println("Error, you need to indicate a host to connect to")
		os.Exit(1)

	}
}
