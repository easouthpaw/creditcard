package creditcard

import (
	"fmt"
	"os"
)

func Run(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Example: creditcard validate [numbers...]")
		os.Exit(1)
	}

	for _, card := range args {
		if LuhnAlgorithm(card) {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			os.Exit(1)
		}
	}
}
