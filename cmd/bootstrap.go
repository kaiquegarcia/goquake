package cmd

import (
	"fmt"
	"os"
)

func Bootstrap() bool {
	if len(os.Args) <= 1 {
		fmt.Println("Missing command. Run 'go run . help' to check the possibilities.")
		return false
	}

	cmd := os.Args[1]

	switch cmd {
	case "report":
		return report()
	case "help":
		help("")
		return false
	}

	return true
}
