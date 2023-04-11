package cmd

import (
	"fmt"
	"goquake/internal/config"
	"os"
)

func Bootstrap() bool {
	args := os.Args[1:]
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
		case "-wap", "--world-as-player":
			if len(args) < index+1 {
				fmt.Println("Missing value for world-as-player argument")
				return false
			}
			index++
			value := args[index]

			config.ConsiderWorldAsPlayer(value == "true" || value == "1")
		case "-gb", "--group-by":
			if len(args) < index+1 {
				fmt.Println("Missing value for group-by argument")
				return false
			}
			index++
			value := args[index]

			config.GroupBy(value)
		case "h", "-h", "help", "--help":
			help()
			return false
		}
	}

	return true
}
