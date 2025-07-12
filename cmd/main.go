package main

import (
	"fmt"
	"github.com/easouthpaw/creditcard"
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
		creditcard.Run(args)
	case "information":
		creditcard.Run(args)
	default:
		fmt.Println("Unknown flag")
		os.Exit(1)
	}
}
