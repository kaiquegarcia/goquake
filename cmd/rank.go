package cmd

import (
	"goquake/internal/config"
	"os"
)

func rank() bool {
	if len(os.Args) == 3 && (os.Args[2] == "-h" || os.Args[2] == "--help") {
		help("rank")
		return false
	}

	config.Ranking(true)
	return true
}
