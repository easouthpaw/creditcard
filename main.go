package main

import (
	"fmt"
	"github.com/easouthpaw/creditcard/cmd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Stderr, "Please use flags validate | test etd.")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "validate":
		cmd.Run(args)
	case "test":
		cmd.Run(args)
	default:
		fmt.Println("Unknown flag")
		os.Exit(1)
	}
}
