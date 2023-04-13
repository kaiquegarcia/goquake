package cmd

import (
	"fmt"
	"goquake/internal/config"
	"os"
)

func rank() bool {
	args := os.Args[2:]
	for index := 0; index < len(args); index++ {
		option := args[index]

		switch option {
		case "-f", "--filepath":
			if len(args) < index+1 {
				fmt.Println("Missing value for filepath argument")
				return false
			}
			index++
			value := args[index]

			config.Filepath(value)
		case "-h", "--help":
			help("rank")
			return false
		}
	}

	config.Ranking(true)
	return true
}
